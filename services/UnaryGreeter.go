package services

import (
	"context"
	"fmt"
	"grpc-tutorial/api/proto/base"
	"grpc-tutorial/api/proto/unary"
)

type UnaryGreeter struct {
	unary.UnimplementedGreeterServer
}

func (u UnaryGreeter) SayHello(ctx context.Context, request *base.HelloRequest) (*base.HelloResponse, error) {
	msg := fmt.Sprintf("Hello, %v!", request.Name)
	fmt.Println(msg)
	return &base.HelloResponse{Message: msg}, nil
}
