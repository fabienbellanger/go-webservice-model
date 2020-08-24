package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// routes list all HTTP routes.
func (s *server) routes() {
	s.router.POST("/post", s.handlePost())
	s.router.GET("/name/:name", s.handleName())
}

func (s *server) handlePost() echo.HandlerFunc {
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

func (s *server) handleName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
	}
}
