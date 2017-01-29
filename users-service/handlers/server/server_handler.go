package handler

import (
	"golang.org/x/net/context"

	pb "github.com/adamryman/ambition-users/users-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UsersServer {
	return usersService{}
}

type usersService struct{}

// CreateUser implements Service.
func (s usersService) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}

// ReadUser implements Service.
func (s usersService) ReadUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}

// UpdateUser implements Service.
func (s usersService) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}

// DeleteUser implements Service.
func (s usersService) DeleteUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}
