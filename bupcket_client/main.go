package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/gnicod/bupcket/bupcket"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "img.png"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUploadClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.UploadRequest{File: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Sendinf file: %s", r.GetUrl())
}
