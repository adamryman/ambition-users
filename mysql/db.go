package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-users/users-service"
)

type database struct {
	db *sql.DB
}

func InitDatabase(conn string) (pb.Database, error) {
	d, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to %s", conn)
	}
	return &database{d}, nil
}

// TODO:
func (d *database) CreateUser(in *pb.User) (*pb.User, error) {
	// TODO: Possibly create userstrello table with trello data and userslocal table for hash, salt, email, etc
	// TODO: Put in all values rather than just trello values
	const query = `INSERT `
	id, err := exec(d.db, query)
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

// TODO:
func (d *database) ReadUserByID(id int64) (*pb.User, error) {
	const query = `SELECT from users where id=?`
	resp := d.db.QueryRow(query, id)
	var user pb.User
	err := resp.Scan( /* user.TrelloInfo. */ )
	if err != nil {
		return nil, err
	}

	// Multi table joins?

	return &user, nil
}

// TODO:
func (d *database) ReadUserByTrelloID(trelloID string) (*pb.User, error) {
	const query = `SELECT from //TODO where trello_id=?`
	resp := d.db.QueryRow(query, trelloID)
	var user pb.User
	err := resp.Scan( /* user.TrelloInfo. */ )
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// exec calls db.db.Exec with passed arguments and returns the id of the LastInsertId
func exec(db *sql.DB, query string, args ...interface{}) (int64, error) {
	resp, err := db.Exec(query, args...)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to exec query: %v", query)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, errors.Wrapf(err, "unable to get last id after query: %v", query)
	}

	return id, nil
}
