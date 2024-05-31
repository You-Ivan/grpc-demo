package main

import (
	"google.golang.org/grpc"
	bi_direction "grpc-tutorial/api/proto/bi-direction"
	"grpc-tutorial/cmd"
	"grpc-tutorial/services"
)

func main() {
	cmd.LaunchService(15004, func(server *grpc.Server) {
		bi_direction.RegisterChatServiceServer(server, &services.MyChatService{})
	})
}
