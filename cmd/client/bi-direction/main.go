package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-tutorial/api/proto/base"
	bi_direction "grpc-tutorial/api/proto/bi-direction"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:15004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := bi_direction.NewChatServiceClient(conn)
	s, err := client.Chat(context.Background())
	for i := 0; i < 10; i++ {
		if err := s.Send(&base.HelloRequest{Name: fmt.Sprintf("Jack %v", i)}); err != nil {
			log.Fatalln(err)
		}
		reply, err := s.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Server: %v\n", reply.Message)
		time.Sleep(time.Second)
	}
	if err := s.CloseSend(); err != nil {
		log.Fatalln(err)
	}
}
