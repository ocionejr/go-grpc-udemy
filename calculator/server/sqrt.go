package main

import (
	"context"
	"log"
	"math"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Number received: %d\n", in.Number)
	number := in.Number


	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}