package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/assultuss/messenger-console/pkg/user_v1"
)

const (
	port = "localhost:50051"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponce, error) {
	log.Printf("Received: %v", in)
	return &pb.CreateUserResponce{Id: 123}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponce, error) {
	log.Printf("Received: %v", in)
	return &pb.GetUserResponce{
		Id:        in.GetId(),
		Name:      "assultuss",
		Email:     "assultuss@gmail.com",
		Role:      pb.Role_USER,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponce, error) {
	log.Printf("Received: %v", in)
	return &pb.UpdateUserResponce{}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponce, error) {
	log.Printf("Received: %v", in)
	return &pb.DeleteUserResponce{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
