package main

import (
	"log"
	"net"
	"os"

	pb "github.com/Punam-Gaikwad/microservices/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":50052"
	defaultHost = "mongodb://localhost:27017"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessels")

	repository := &VesselRepository{vesselCollection}

	pb.RegisterVesselServiceServer(s, &handler{repository})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
