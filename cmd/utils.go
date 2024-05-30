package cmd

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// ServerFunc defines a function that registers a gRPC service to a server.
type ServerFunc func(server *grpc.Server)

func LaunchService(port int, registerService ServerFunc) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	registerService(server)

	fmt.Printf("Server is listening on port: %d\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
