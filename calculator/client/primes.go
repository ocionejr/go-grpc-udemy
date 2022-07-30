package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling doPrime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Result: %d\n", msg.Result)
	}
}