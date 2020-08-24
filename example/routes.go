package example

import "github.com/labstack/echo/v4"

// routes list all HTTP routes.
func (s *Server) routes() {
	s.Router.GET("/example", s.handleExample())
}

func (s *Server) handleExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		s.Store.getExamples()

		t := make([]Example, 0)
		for i := 0; i < 100000; i++ {
			t = append(t, Example{
				ID:      int64(i),
				Name:    "Coucou ceci est mon nom",
				Message: "Mon message doit être un peu long pour augmenter la taille",
			})
		}

		return c.JSON(200, t)
	}
}
