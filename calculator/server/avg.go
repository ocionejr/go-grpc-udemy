package main

import (
	"io"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	sumNumbers := float64(0)
	totalNumbers := float64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: sumNumbers / totalNumbers,
			})
		}

		if err != nil {
			log.Fatalf("Error while receiving stream from client: %v\n", err)
		}

		sumNumbers += float64(req.Number)
		totalNumbers++
	}
}