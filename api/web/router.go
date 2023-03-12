package web

import (
	"doc-api/api/controller"
	"doc-api/env"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type (
	Router interface {
		RegisterEndpoint(e *echo.Echo)
	}

	router struct {
		// env
		secret string
		// middleware
		authUser controller.AuthUserMiddleware
		// endpoint
		helloHandler controller.HelloHandler
		loginHandler controller.LoginHandler
		docHandler   controller.DocHandler
		userHandler  controller.UserHandler
	}
)

func NewRouter(
	secret env.JWT_SECRET_KEY, // env
	authUser controller.AuthUserMiddleware, // middleware
	helloHandler controller.HelloHandler, // endpoint
	loginHandler controller.LoginHandler, // endpoint
	docHandler controller.DocHandler, //endpoint
	userHandler controller.UserHandler, //endpoint
) Router {
	return &router{
		secret:       string(secret),
		authUser:     authUser,
		helloHandler: helloHandler,
		loginHandler: loginHandler,
		docHandler:   docHandler,
		userHandler:  userHandler,
	}
}

func (r *router) RegisterEndpoint(e *echo.Echo) {
	// routes
	// free access
	e.GET("/", r.helloHandler.Get)
	e.POST("/signup", r.userHandler.Post)
	e.POST("/login", r.loginHandler.Post)

	// restricted access(role user)
	user := e.Group("/user")
	user.Use(echojwt.WithConfig(echojwt.Config{SigningKey: []byte(r.secret)})) // jwt valid
	user.Use(r.authUser.AuthUser())                                            // user_id validation
	user.GET("/", r.userHandler.Get)
	user.PUT("/", r.userHandler.Put)
	user.DELETE("/", r.userHandler.Delete)
	user.GET("/docs", r.docHandler.List)
	user.GET("/docs/:id", r.docHandler.Get)
	user.POST("/docs", r.docHandler.Post)
	user.PUT("/docs/:id", r.docHandler.Put)
	user.DELETE("/docs/:id", r.docHandler.Delete)

	// restricted access(role adminonly)
	admin := e.Group("/admin")
	admin.Use(echojwt.WithConfig(echojwt.Config{SigningKey: []byte(r.secret)})) // jwt valid
	admin.Use(r.authUser.AuthAdminUser())                                       // user_id validation
}
