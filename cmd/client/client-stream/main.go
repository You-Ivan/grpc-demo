package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	clientstream "grpc-tutorial/api/proto/client-stream"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:15003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	client := clientstream.NewTemperatureServiceClient(conn)

	client, err := client.GetAvgTemp(context.Background())
	for i := 0; i < 10; i++ {
		err := client.Send(&clientstream.Temperature{Time: timestamppb.Now(), Degree: float64(i)})
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
	_, err = client.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
}
