package mysql

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-users/users-service"
)

type Database struct {
	db *sql.DB
}

func Open(conn string) (*Database, error) {
	d, err := sql.Open("sqlite3", conn)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to %s", conn)
	}
	if err := d.Ping(); err != nil {
		return nil, err
	}
	if err := setupDB(d); err != nil {
		return nil, err
	}
	return &Database{d}, nil
}

func setupDB(db *sql.DB) error {
	const usersTable = `CREATE TABLE IF NOT EXISTS users(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username text,
				email text DEFAULT '',
				trello_id text DEFAULT '',
				trello_webhookurl text DEFAULT '')`
	_, err := db.Exec(usersTable)
	if err != nil {
		return err
	}

	const adam = `DELETE FROM users WHERE id = 1; INSERT INTO users (id, username, trello_id) VALUES (?, ?, ?)`
	_, err = db.Exec(adam, 1, "adamryman", os.Getenv("TRELLO_ID"))
	if err != nil {
		return err
	}

	return nil
}

// TODO:
func (d *Database) CreateUser(in *pb.User) (*pb.User, error) {
	// TODO: Possibly create userstrello table with trello data and userslocal table for hash, salt, email, etc
	// TODO: Put in all values rather than just trello values
	const query = `INSERT INTO users (username, email, trello_id, trello_webhookurl) VALUES(?, ?, ?, ?)`
	id, err := exec(d.db, query, in.Info.Username, in.Info.Email, in.Trello.ID, in.Trello.WebhookURL)
	if err != nil {
		return nil, err
	}
	in.ID = id

	return in, nil
}

func (d *Database) ReadUserByID(id int64) (*pb.User, error) {
	const query = `SELECT * FROM users WHERE id=?`
	resp := d.db.QueryRow(query, id)
	var user pb.User
	user.Trello = &pb.TrelloInfo{}
	user.Info = &pb.UserInfo{}
	err := resp.Scan(&user.ID, &user.Info.Username, &user.Info.Email, &user.Trello.ID, &user.Trello.WebhookURL)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) ReadUserByTrelloID(trelloID string) (*pb.User, error) {
	const query = `SELECT * FROM users WHERE trello_id=?`
	resp := d.db.QueryRow(query, trelloID)
	var user pb.User
	user.Trello = &pb.TrelloInfo{}
	user.Info = &pb.UserInfo{}
	err := resp.Scan(&user.ID, &user.Info.Username, &user.Info.Email, &user.Trello.ID, &user.Trello.WebhookURL)
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
