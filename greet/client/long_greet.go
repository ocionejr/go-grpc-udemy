package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ocionejr/go-grpc-udemy/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Ocione"},
		{FirstName: "Leo"},
		{FirstName: "Vini"},
		{FirstName: "Salva"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving error from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
