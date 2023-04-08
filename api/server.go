package api

import (
	"doc-api/api/web"
	"doc-api/env"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Server struct {
		PORT          string
		router        web.Router
		loggerFactory web.LoggerFactory
	}
)

func NewServer(port env.PORT, router web.Router, loggerFactory web.LoggerFactory) Server {
	return Server{
		PORT:          string(port),
		router:        router,
		loggerFactory: loggerFactory,
	}
}

func (s *Server) Run() {
	e := echo.New()
	log.SetFlags(0)

	e.Use(s.loggerFactory.Middleware())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	// routes
	s.router.RegisterEndpoint(e)

	e.Logger.Fatal(e.Start(":" + s.PORT))
}
