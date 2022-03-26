package service

import (
	"context"
	pb "grpcex/proto"
)

// server is used to implement helloworld.GreeterServer.
type ExampleService struct {
	pb.UnimplementedGreeterServer
}

func NewService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return nil, nil
}
