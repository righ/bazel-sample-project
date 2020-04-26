package main

import (
	"context"
	"log"
	"net"

	pb "github.com/righ/go-sample-bazel-project/protobuf"

	"google.golang.org/grpc"
)

var conn pb.EchoClient

type message struct {
	Message string
}

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) Echo(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return in, nil
}

func main() {
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
