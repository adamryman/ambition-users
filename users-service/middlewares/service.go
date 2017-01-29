package middlewares

import (
	pb "github.com/adamryman/ambition-users/users-service"
)

func WrapService(in pb.UsersServer) pb.UsersServer {
	return in
}
