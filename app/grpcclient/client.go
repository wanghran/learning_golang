package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

const (
	host        = "127.0.0.1"
	defaultName = "world"
)

func main() {
	address := fmt.Sprintf("%s:%v", host, common.Port)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

}
