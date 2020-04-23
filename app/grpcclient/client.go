package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/wanghran/learning_golang/api/protobuf-spec"
	"google.golang.org/grpc"
)

const (
	host        = "127.0.0.1"
	defaultName = "world"
)

func main() {
	address := fmt.Sprintf("%s:%v", host, 9090)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
