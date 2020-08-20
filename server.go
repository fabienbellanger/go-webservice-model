package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// server represents
type server struct {
	store  Store
	router *echo.Echo
}

// newServer creates a new instance of Server.
func newServer() *server {
	s := &server{
		router: echo.New(),
	}

	s.routes()

	return s
}

func (s *server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
