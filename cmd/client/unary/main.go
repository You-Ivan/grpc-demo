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
	//create a connection to the server
	conn, err := grpc.NewClient("localhost:15001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	// create a stub of the server
	client := unary.NewGreeterClient(conn)
	// invoke the server method
	resp, err := client.SayHello(context.Background(), &base.HelloRequest{Name: "Jack"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Message)

}
