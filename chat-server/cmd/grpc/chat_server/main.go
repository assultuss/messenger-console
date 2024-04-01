package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/assultuss/messenger-console/pkg/chat_v1"
)

const port = ":50052"

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateChatRequest) (*pb.CreateChatResponce, error) {
	log.Printf("Received: %v", in.GetUsernames())
	return &pb.CreateChatResponce{Id: 123}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteChatRequest) (*pb.DeleteChatResponce, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.DeleteChatResponce{}, nil
}

func (s *server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponce, error) {
	log.Printf("Received: %v", in)
	return &pb.SendMessageResponce{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
