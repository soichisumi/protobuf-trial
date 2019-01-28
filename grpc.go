package main

import (
	"context"
	"fmt"
	"github.com/soichisumi/protobuf-trial/pbtrial"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Users map[string]pbtrial.User //todo: goroutine unsafe
}

func NewServer() Server {
	userDB := make(map[string]pbtrial.User)
	return Server{
		Users: userDB,
	}
}

func (s *Server) AddUser(ctx context.Context, req *pbtrial.AddUserRequest) (*pbtrial.AddUserResponse, error) {
	user := req.User
	if user.Name == "" {
		fmt.Printf("username is empty. user: %+v\n", user)
		return &pbtrial.AddUserResponse{}, status.Error(codes.InvalidArgument, "")
	}
	s.Users[user.Name] = *user

	return &pbtrial.AddUserResponse{}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pbtrial.GetUserRequest) (*pbtrial.GetUserResponse, error) {
	username := req.Username
	if username == "" {
		fmt.Printf("username is empty. username: %s\n", username)
		return &pbtrial.GetUserResponse{}, status.Error(codes.InvalidArgument, "")
	}
	user, ok := s.Users[username]
	if !ok {
		fmt.Println("given user is not found")
		return &pbtrial.GetUserResponse{}, status.Error(codes.NotFound, "")
	}
	return &pbtrial.GetUserResponse{
		User: &user,
	}, nil
}