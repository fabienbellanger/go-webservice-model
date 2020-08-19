package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

// Store is an interface that represents the store of the application.
type Store interface {
	Open() error
	Close() error

	GetSuperUsers() ([]*SuperUser, error)
}

type dbStore struct {
	db *sqlx.DB
}

func (s *dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/pos")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	s.db = db

	return nil
}

func (s *dbStore) Close() error {
	return s.db.Close()
}

func (s *dbStore) GetSuperUsers() ([]*SuperUser, error) {
	var users []*SuperUser

	err := s.db.Select(&users, "SELECT * FROM SuperUser")
	if err != nil {
		return users, err
	}
	return users, nil
}
