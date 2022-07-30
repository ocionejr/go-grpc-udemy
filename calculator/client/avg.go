package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ocionejr/go-grpc-udemy/calculator/proto"
)


func doAvg(c pb.CalculatorServiceClient) {
	numbers := []int32 {1,2,3,4}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg: %s\n", err)
	}

	for _, num := range numbers {
		stream.Send(&pb.AvgRequest{Number: int64(num)})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	
	if err != nil {
		log.Fatalf("Error while receiving from Avg: %s\n", err)
	}

	log.Printf("Res: %g", res.Avg)
}