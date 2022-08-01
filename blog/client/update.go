package main

import (
	"context"
	"log"

	pb "github.com/ocionejr/go-grpc-udemy/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string){
	log.Println("Update blog was invoked")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Not Ocione",
		Title: "A new title",
		Content: "Content of the first blog with some awesome additions",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error while updating: %v", err)
	}

	log.Println("Blog was updated!")
}