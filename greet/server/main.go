package main

import (
	"log"
	"net"

	pb "github.com/rajesh4b8/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

var (
	addr = "127.0.0.1:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
