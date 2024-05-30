package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-tutorial/api/proto/base"
	"grpc-tutorial/api/proto/unary"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:15001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := unary.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &base.HelloRequest{Name: "Jack"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Message)

}
