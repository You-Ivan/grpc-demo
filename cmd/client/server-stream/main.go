package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-tutorial/api/proto/base"
	server_stream "grpc-tutorial/api/proto/server-stream"
	"io"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:15002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client := server_stream.NewGreeterClient(conn)
	stream, err := client.SayHelloStream(context.Background(), &base.HelloRequest{Name: "Ivan"})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}
		log.Printf("Received response: %v", response)
	}
}
