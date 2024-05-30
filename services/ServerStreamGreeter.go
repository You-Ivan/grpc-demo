package services

import (
	"fmt"
	"grpc-tutorial/api/proto/base"
	serverstream "grpc-tutorial/api/proto/server-stream"
	"log"
	"time"
)

type ServerStreamGreeter struct {
	serverstream.UnimplementedGreeterServer
}

func (s ServerStreamGreeter) SayHelloStream(request *base.HelloRequest, server serverstream.Greeter_SayHelloStreamServer) error {
	name := request.Name
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		err := server.Send(&base.HelloResponse{
			Message: fmt.Sprintf("Hello, %v!", name),
		})
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
