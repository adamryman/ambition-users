package handler

import (
	"golang.org/x/net/context"

	"github.com/adamryman/dbconn"
	"github.com/pkg/errors"

	"github.com/adamryman/ambition-users/mysql"
	pb "github.com/adamryman/ambition-users/users-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UsersServer {
	database, err := mysql.InitDatabase(dbconn.FromENV().MySQL())
	if err != nil {
		panic(err)
	}
	return usersService{
		db: database,
	}
}

type usersService struct {
	db pb.Database
}

// CreateUser implements Service.
// TODO:
func (s usersService) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	// TODO: input validation
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
	if id := in.GetID(); id != 0 {
		return s.db.ReadUserByID(id)
	}
	if id := in.GetTrello().GetID(); id != "" {
		return s.db.ReadUserByTrelloID(id)
	}
	return nil, errors.New("cannot read action, need ID or TrelloInfo.ID")
}

// UpdateUser implements Service.
// TODO: non-MVP
func (s usersService) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	// TODO: input validation
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}

// DeleteUser implements Service.
// TODO: non-MVP
func (s usersService) DeleteUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	// TODO: input validation
	var resp pb.User
	resp = pb.User{
	// ID:
	// Info:
	// Trello:
	}
	return &resp, nil
}
