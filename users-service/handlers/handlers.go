package handlers

import (
	"golang.org/x/net/context"
	"os"

	"github.com/pkg/errors"

	sql "github.com/adamryman/ambition-users/sqlite"
	pb "github.com/adamryman/ambition-users/users-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UsersServer {
	database, err := sql.Open(os.Getenv("SQLITE3"))
	if err != nil {
		// TODO: Do not panic, start something to try connection over and over.
		// Maybe 100 times?
		// DEBUG_SVC=1 then do like 3.
		// There will also need to be retry logic for the database methods
		panic(err)
	}
	return usersService{
		db: database,
	}
}

type usersService struct {
	db *sql.Database
}

// CreateUser implements Service.
// TODO:
func (s usersService) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	// TODO: input validation
	u, err := s.db.CreateUser(in)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create user")
	}
	return u, nil
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
