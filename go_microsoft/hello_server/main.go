package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// wrappedStream 包装 grpc.ServerStream，用于在流式拦截器中传递自定义上下文
type wrappedStream struct {
	grpc.ServerStream
}

// 实现 grpc.ServerStream 接口的 Context 方法
// 这样可以返回自定义的上下文，而不是原始的上下文
func (w *wrappedStream) Context() context.Context {
	// 可以返回带有额外值的上下文
	return w.ServerStream.Context()
}

// newWrappedStream 创建包装后的流
func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token认证的逻辑
	// 这里是为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}

// unaryInterceptor 服务端一元拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	m, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return m, err
}

// streamInterceptor 服务端流拦截器
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return err
}

// grpc server
type server struct {
	pb.UnimplementedGreeterServer                //没有实现方法也可以跑起来
	count                         map[string]int //记录次数
	mu                            sync.Mutex     //并发锁
}

// Sayhello是我们的业务逻辑
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//在执行业务逻辑之前要check metadata是否包含token
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.GetName()]++ //记录次数
	if s.count[in.Name] > 1 {
		//返回请求次数限制的错误
		st := status.New(codes.ResourceExhausted, "request limit")
		//添加错误的详情信息
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "每个name只能调用一次SayHello",
				},
				},
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}
	vl := md.Get("token")
	if vl[0] != "app-test-jasper" {
		return nil, status.Error(codes.Unauthenticated, "无效toekn")
	}
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	//启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen,err:%v\n", err)
		return
	}
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed err:%v", creds)
	}
	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	//创建grpc服务
	//注册服务
	pb.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	//启动服务
	// err = s.Serve(l)
	//开启goroutine
	go func() {
		log.Fatalln(s.Serve(l))
	}()
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
	}
	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.NewClient(
		"localhost:8972",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8072",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8072")
	log.Fatalln(gwServer.ListenAndServe())
}
