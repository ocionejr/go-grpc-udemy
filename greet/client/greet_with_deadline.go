package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ocionejr/go-grpc-udemy/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest {
		FirstName: "Ocione",
	}

	res, err := c.GreetWithDealine(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			}else{
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}

		} else {
			log.Fatalf("Non gRPC erorr: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}