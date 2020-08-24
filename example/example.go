package example

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
	Store  Store
}

func NewServer(router *echo.Echo) *Server {
	s := &Server{
		Router: router,
	}

	s.routes()

	return s
}
