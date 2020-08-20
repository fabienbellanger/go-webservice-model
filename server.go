package gwm

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	Store  Store
	Router *echo.Echo
}

// NewServer creates a new instance of Server.
func NewServer() *Server {
	s := &Server{
		Router: echo.New(),
	}

	s.initHTTPServer()
	s.routes()

	return s
}

func (s *Server) initHTTPServer() {
	// Startup banner
	// --------------
	if viper.GetString("environment") == "production" {
		s.Router.HideBanner = true
	}

	// Logger
	// ------
	s.Router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} | ${remote_ip}\t| ${status} | ${method} | ${uri} | ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           os.Stderr,
	}))

	// Recover
	// -------
	s.Router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisableStackAll:   true,
		DisablePrintStack: true,
		Skipper:           middleware.DefaultSkipper,
	}))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
