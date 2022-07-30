package main

import (
	"context"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func (s *Server) SumNumbers(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Response: in.FirstNumber + in.SecondNumber,
	}, nil
}