package example

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

// routes list all HTTP routes.
func (s *Server) Routes() {
	s.Router.GET("/example", s.handleExample())
}

func (s *Server) handleExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "Hello example!")
	}
}
