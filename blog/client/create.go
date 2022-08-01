package main

import (
	"context"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Ocione",
		Title: "My first blog",
		Content: "Content",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has beeen created: %s\n", res.Id)
	return res.Id
}