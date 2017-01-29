package clienthandler

import (
	pb "github.com/adamryman/ambition-users/users-service"
)

// CreateUser implements Service.
func CreateUser(IDCreateUser int64, InfoCreateUser pb.UserInfo, TrelloCreateUser pb.TrelloInfo) (*pb.User, error) {
	request := pb.User{
		ID:     IDCreateUser,
		Info:   &InfoCreateUser,
		Trello: &TrelloCreateUser,
	}
	return &request, nil
}

// ReadUser implements Service.
func ReadUser(IDReadUser int64, InfoReadUser pb.UserInfo, TrelloReadUser pb.TrelloInfo) (*pb.User, error) {
	request := pb.User{
		ID:     IDReadUser,
		Info:   &InfoReadUser,
		Trello: &TrelloReadUser,
	}
	return &request, nil
}

// UpdateUser implements Service.
func UpdateUser(IDUpdateUser int64, InfoUpdateUser pb.UserInfo, TrelloUpdateUser pb.TrelloInfo) (*pb.User, error) {
	request := pb.User{
		ID:     IDUpdateUser,
		Info:   &InfoUpdateUser,
		Trello: &TrelloUpdateUser,
	}
	return &request, nil
}

// DeleteUser implements Service.
func DeleteUser(IDDeleteUser int64, InfoDeleteUser pb.UserInfo, TrelloDeleteUser pb.TrelloInfo) (*pb.User, error) {
	request := pb.User{
		ID:     IDDeleteUser,
		Info:   &InfoDeleteUser,
		Trello: &TrelloDeleteUser,
	}
	return &request, nil
}
