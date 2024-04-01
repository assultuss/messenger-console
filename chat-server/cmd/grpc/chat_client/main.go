package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/assultuss/messenger-console/pkg/chat_v1"
)

const address = "localhost:50052"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	createChat(client)
	deleteChat(client)
	sendMessage(client)
}

func createChat(client pb.ChatServiceClient) {
	req := &pb.CreateChatRequest{
		Usernames: []string{"user1", "user2"},
	}
	res, err := client.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating chat: %v", err)
	}
	fmt.Println("Created chat with ID: %d\n", res.Id)
}

func deleteChat(client pb.ChatServiceClient) {
	req := &pb.DeleteChatRequest{
		Id: 123,
	}
	_, err := client.Delete(context.Background(), req)
	if err != nil {
		log.Fatalf("Error deleting chat: %v", err)
	}
	fmt.Println("Deleted chat successfully")
}

func sendMessage(client pb.ChatServiceClient) {
	req := &pb.SendMessageRequest{
		From:      "user1",
		Text:      "Hello, world!",
		Timestamp: nil,
	}
	_, err := client.SendMessage(context.Background(), req)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	fmt.Println("Message sent successfully")
}
