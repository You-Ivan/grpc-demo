package main

import (
	"google.golang.org/grpc"
	"grpc-tutorial/api/proto/unary"
	"grpc-tutorial/cmd"
	"grpc-tutorial/services"
)

func main() {
	f := func(server *grpc.Server) {
		unary.RegisterGreeterServer(server, services.UnaryGreeter{})
	}
	cmd.LaunchService(15001, f)
}
