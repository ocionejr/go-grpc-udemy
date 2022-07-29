package main

import (
	"context"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ocione",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

func doSum(c pb.GreetServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.SumNumbers(context.Background(), &pb.SumRequest{
		FirstNumber: 12,
		SecondNumber: 4,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}