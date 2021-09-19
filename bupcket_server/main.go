package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gnicod/bupcket/bupcket"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUploadServer
}

func (s *server) SayHello(ctx context.Context, in *pb.UploadRequest) (*pb.UploadReply, error) {
	log.Printf("Should upload: %v", in.GetFile())
	return &pb.UploadReply{Url: "http://fakeurl.co/" + in.GetFile()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUploadServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}