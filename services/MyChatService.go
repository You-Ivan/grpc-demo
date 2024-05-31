package services

import (
	"fmt"
	"grpc-tutorial/api/proto/base"
	"grpc-tutorial/api/proto/bi-direction"
	"io"
	"log"
)

type MyChatService struct {
	bi_direction.UnimplementedChatServiceServer
}

func (m MyChatService) Chat(stream bi_direction.ChatService_ChatServer) error {
	for {
		helloReq, err := stream.Recv()
		if err == io.EOF {
			log.Println("Chat Ended")
			return nil
		}
		name := helloReq.Name
		log.Printf("Received a message from: %v", name)
		err = stream.Send(&base.HelloResponse{Message: fmt.Sprintf("Hello, %v!", name)})
		if err != nil {
			return err
		}
	}
}
