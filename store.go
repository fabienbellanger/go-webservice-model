package gwm

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

// Store is an interface that represents the store of the application.
type Store interface {
	Open() error
	Close() error

	GetSuperUsers() ([]*SuperUser, error)
}

type DBStore struct {
	DB *sqlx.DB
}

// Open opens database connection.
func (s *DBStore) Open() error {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name")))
	if err != nil {
		return err
	}
	log.Printf("Connected to DB %s", viper.GetString("database.name"))
	s.DB = db

	return nil
}

// Close closes database conection.
func (s *DBStore) Close() error {
	return s.DB.Close()
}

// GetSuperUsers returns SuperUsers list.
func (s *DBStore) GetSuperUsers() ([]*SuperUser, error) {
	var users []*SuperUser

	err := s.DB.Select(&users, "SELECT * FROM SuperUser")
	if err != nil {
		return users, err
	}
	return users, nil
}
