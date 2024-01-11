package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/hello-server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc服务端中去注册我们的自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("grpcServer Start Filed")
	}
}
