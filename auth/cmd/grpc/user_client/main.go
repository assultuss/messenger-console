package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/assultuss/messenger-console/pkg/user_v1"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	createUser(client)
	getUser(client)
	updateUser(client)
	deleteUser(client)
}

func createUser(client pb.UserServiceClient) {
	req := &pb.CreateUserRequest{
		Name:            "giga",
		Email:           "kek@228.giga",
		Password:        "password1234",
		PasswordConfirm: "password1234",
		Role:            pb.Role_USER,
	}

	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("Created user with ID: %d\n", res.Id)
}

func getUser(client pb.UserServiceClient) {
	req := &pb.GetUserRequest{
		Id: 123,
	}

	res, err := client.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	fmt.Printf("Got user: %+v\n", res)
}

func updateUser(client pb.UserServiceClient) {
	req := &pb.UpdateUserRequest{
		Id:    123,
		Name:  "assultuss",
		Email: "assultuss@gmail.cock",
	}

	_, err := client.UpdateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("Updated user successfully")
}

func deleteUser(client pb.UserServiceClient) {
	req := &pb.DeleteUserRequest{
		Id: 123,
	}

	_, err := client.DeleteUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	fmt.Println("Deleted user successfully")
}
