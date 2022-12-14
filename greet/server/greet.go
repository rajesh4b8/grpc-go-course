package main

import (
	"context"
	"log"

	pb "github.com/rajesh4b8/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	// business logic...

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
