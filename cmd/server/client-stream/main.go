package main

import (
	"google.golang.org/grpc"
	client_stream "grpc-tutorial/api/proto/client-stream"
	"grpc-tutorial/cmd"
	"grpc-tutorial/services"
)

func main() {
	cmd.LaunchService(15003, func(server *grpc.Server) {
		client_stream.RegisterTemperatureServiceServer(server, services.MyTemperatureService{})
	})
}
