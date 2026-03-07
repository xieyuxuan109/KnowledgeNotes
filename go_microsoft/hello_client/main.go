package main

import (
	"context"
	"flag"
	"fmt"
	"hello_client/pb"
	"io"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

type wrappedStream struct {
	grpc.ClientStream
}

// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// streamInterceptor 客户端流式拦截器
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

// grpc客户端
// 调用server端SayHello方法
func main() {
	flag.Parse() //解析命令行参数
	//记载证书
	creds, _ := credentials.NewClientTLSFromFile("certs/server.crt", "")
	//连接server
	conn, err := grpc.NewClient("127.0.0.1:8972",
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithStreamInterceptor(streamInterceptor),
	)

	if err != nil {
		log.Fatalf("grpc.Dail failed,err:%v", err)
	}
	defer conn.Close()
	//创建客户端
	c := pb.NewGreeterClient(conn) //自动生成的代码
	//调用rpc方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//带元数据
	md := metadata.Pairs(
		"token", "app-test-jasper",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("QuotaFailure:%s\n", info)
			default:
				fmt.Printf("unexpected type:%s\n", info)
			}
		}
		log.Printf("c.SayHello failed, err%v", err)
		return
	}

	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("c.LotsOfReplies failed, err%v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv failed, err:%v\n", err)
			return
		}
		log.Printf("recv:%v\n", res.GetReply())
	}
	//拿到rpc响应
	log.Printf("resp:%v", resp.GetReply())
}
