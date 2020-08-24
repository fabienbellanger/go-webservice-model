package example

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Init(db *sqlx.DB)

	getExamples() error
}

type DBStore struct {
	DB *sqlx.DB
}

func (s *DBStore) Init(db *sqlx.DB) {
	s.DB = db
}

func (s *DBStore) getExamples() error {
	log.Println("GetExamples")
	return nil
}
