package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"features/data"
	ecpb "features/proto/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}

func main() {
	flag.Parse()

	// 客户端通过ca证书来验证服务的提供的证书
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca.crt"), "www.xiaomingdou.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	// 建立连接时指定使用 TLS
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := ecpb.NewEchoClient(conn)
	callUnaryEcho(client, "hello world")
}
