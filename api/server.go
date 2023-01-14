package api

import (
	"doc-api/env"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Server struct {
		PORT string
	}
)

func NewServer(port env.PORT) Server {
	return Server{
		PORT: string(port),
	}
}

func (s *Server) Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Hello world!") })

	e.Logger.Fatal(e.Start(":" + s.PORT))
}
