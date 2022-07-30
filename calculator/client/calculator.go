package main

import (
	"context"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.SumNumbers(context.Background(), &pb.SumRequest{
		FirstNumber: 12,
		SecondNumber: 4,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Response)
}