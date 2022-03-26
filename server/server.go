package server

import (
	"fmt"
	"log"
	"net"

	pb "grpcex/proto"
	"grpcex/service"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	port = "5001"
)

func NewServer() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service := service.NewService()

	pb.RegisterGreeterServer(s, service)
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
