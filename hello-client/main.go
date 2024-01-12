package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/hello-server/proto"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appID":  "YJZ",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("D:\\GoLandProject\\grpc-demo\\key\\test.pem",
	//	"*.kuangstudy.com")
	//链接到server段，此处禁用安全传输，没有加密和验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not connnect: %v", err)
	}
	defer conn.Close()

	//建立链接
	client := pb.NewSayHelloClient(conn)

	//执行rpc调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{
		RequestName: "yjz",
	})

	fmt.Println(resp.GetResponseMsg())
}
