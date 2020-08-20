package gwm

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// routes list all HTTP routes.
func (s *Server) routes() {
	s.Router.POST("/post", s.handlePost())
	s.Router.GET("/name/:name", s.handleName())
}

func (s *Server) handlePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		l := struct {
			Login string `json:"login" form:"login" query:"login"`
		}{}
		log.Printf("%v", &l)

		if err := c.Bind(&l); err != nil {
			return nil
		}
		return c.JSON(http.StatusOK, l)
	}
}

func (s *Server) handleName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
	}
}
