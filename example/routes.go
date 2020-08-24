package example

import "github.com/labstack/echo/v4"

// routes list all HTTP routes.
func (s *Server) routes() {
	s.Router.GET("/example", s.handleExample())
}

func (s *Server) handleExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		s.Store.getExamples()
		return c.String(200, "Hello example!")
	}
}
