package api

import (
	"doc-api/api/web"
	"doc-api/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Server struct {
		PORT   string
		router web.Router
	}
)

func NewServer(port env.PORT, router web.Router) Server {
	return Server{
		PORT:   string(port),
		router: router,
	}
}

func (s *Server) Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	s.router.RegisterEndpoint(e)

	e.Logger.Fatal(e.Start(":" + s.PORT))
}
