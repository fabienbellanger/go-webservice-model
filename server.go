package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type server struct {
	store  store
	router *echo.Echo
}

// newServer creates a new instance of Server.
func newServer() *server {
	s := &server{
		router: echo.New(),
	}

	s.initHTTPServer()
	s.routes()

	return s
}

func (s *server) initHTTPServer() {
	// Startup banner
	// --------------
	if viper.GetString("environment") == "production" {
		s.router.HideBanner = true
	}

	// Logger
	// ------
	s.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} | ${remote_ip}\t| ${status} | ${method} | ${uri} | ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           os.Stdout,
	}))

	// Recover
	// -------
	s.router.Use(middleware.Recover())
}

func (s *server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
