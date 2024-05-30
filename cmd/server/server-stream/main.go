package main

import (
	"google.golang.org/grpc"
	server_stream "grpc-tutorial/api/proto/server-stream"
	"grpc-tutorial/cmd"
	"grpc-tutorial/services"
)

func main() {
	f := func(server *grpc.Server) {
		server_stream.RegisterGreeterServer(server, services.ServerStreamGreeter{})
	}
	cmd.LaunchService(15002, f)
}
