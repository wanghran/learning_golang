package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wanghran/learning_golang/api/protobuf-spec/grpc_example"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %d with error: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server gRPC with error: %v", err)
	}

}
