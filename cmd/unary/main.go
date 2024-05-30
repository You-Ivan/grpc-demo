package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-tutorial/api/proto/unary"
	"grpc-tutorial/services"
	"log"
	"net"
)

var (
	port = 50051
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	unary.RegisterGreeterServer(server, &services.UnaryGreeter{})
	err = server.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server is listening on port:")
}
