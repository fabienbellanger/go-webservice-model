package example

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Store interface {
	Init(db *sqlx.DB) error

	GetExamples() error
}

type Server struct {
	Router *echo.Echo
	Store  Store
}

type DBStore struct {
	DB *sqlx.DB
}

func NewServer(router *echo.Echo) *Server {
	s := &Server{
		Router: router,
	}

	s.routes()

	return s
}

func (s *DBStore) Init(db *sqlx.DB) error {
	s.DB = db
	return nil
}

// routes list all HTTP routes.
func (s *Server) routes() {
	s.Router.GET("/example", s.handleExample())
}

func (s *Server) handleExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		s.Store.GetExamples()
		return c.String(200, "Hello example!")
	}
}

func (s *DBStore) GetExamples() error {
	log.Println("GetExamples")
	return nil
}
