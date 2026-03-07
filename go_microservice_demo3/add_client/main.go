package main

import (
	"add_client/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//连接rpc server
	conn, err := grpc.NewClient("127.0.0.1:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient failed err:%v", err)
	}
	defer conn.Close()
	//创建rpc client客户端
	client := proto.NewCalcServiceClient(conn)
	//发起rpc调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{X: 10, Y: 20})
	if err != nil {
		log.Fatalf("client.Add failed err:%v", err)
	}
	//打印结果
	log.Printf("resp:%v", resp.GetResult())
}
