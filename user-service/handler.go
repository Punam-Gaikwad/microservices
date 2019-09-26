package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/Punam-Gaikwad/microservices/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenservice Authable
}

func (s *service) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Response{User: user}, nil
}

func (s *service) GetAll(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.Response{Users: users}, nil
}

func (s *service) Auth(ctx context.Context, req *pb.User) (*pb.Token, error) {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := s.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return nil, err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := s.tokenservice.Encode(user)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (s *service) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {
	// Generates a hashed version of our password
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if _, err := s.repo.Create(req); err != nil {
		return nil, errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := s.tokenservice.Encode(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: req, Token: &pb.Token{Token: token}}, nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	// Decode token
	claims, err := s.tokenservice.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	if claims.User.Id == "" {
		return nil, errors.New("invalid user")
	}
	return &pb.Token{Valid: true}, nil
}
