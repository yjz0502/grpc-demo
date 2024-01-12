package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "grpc-demo/hello-server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

// 业务
func (server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	//获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx) //md里面的所有字符都会变成小写
	if !ok {
		return nil, errors.New("未传输 token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "YJZ" || appKey != "123123" {
		return nil, errors.New("token 不正确")
	}
	fmt.Println("hello ", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	//creds, _ := credentials.NewServerTLSFromFile("D:\\GoLandProject\\grpc-demo\\key\\test.pem",
	//	"D:\\GoLandProject\\grpc-demo\\key\\test.key")

	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//在grpc服务端中去注册我们的自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpcServer Start Filed")
		return
	}
}
