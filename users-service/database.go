package users

// Input to database must have necessary fields populated
type Database interface {
	// Requires nothing, though will be very unhelpful if nothing is providied
	CreateUser(in *User) (*User, error)

	ReadUserByID(id int64) (*User, error)
	ReadUserByTrelloID(trelloID int64) (*User, error)
	// TODO: Non-MVP
}
