package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Punam-Gaikwad/microservices/user-service/proto/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50053"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	fmt.Println("user service on address: ", address)

	users := []*pb.User{
		&pb.User{Name: "name", Email: "email", Password: "password", Company: "company"},
	}

	for _, v := range users {
		res, err := client.Create(context.Background(), v)
		if err != nil {
			log.Fatalf("Could not greet: %v", err)
		}
		log.Printf("User Created:\n %v", res.User)
	}
}
