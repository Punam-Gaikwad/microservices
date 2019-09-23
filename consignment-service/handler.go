package main

import (
	"fmt"

	pb "github.com/Punam-Gaikwad/microservices/consignment-service/proto/consignment"
	vesselProto "github.com/Punam-Gaikwad/microservices/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

//Handler -
type Handler struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save our consignment
	er := s.repo.Create(req)
	//consignment,err := s.repo.Create(req)
	if er != nil {
		return nil, er
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true}, nil
}

func (s *Handler) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	fmt.Println("entered to GetAll service method")

	consignments, _ := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}
