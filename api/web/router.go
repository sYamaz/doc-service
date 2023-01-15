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
		helloHandler  controller.HelloHandler
		loginHandler  controller.LoginHandler
		signupHandler controller.SignupController
	}
)

func NewRouter(
	secret env.JWT_SECRET_KEY, // env
	authUser controller.AuthUserMiddleware, // middleware
	helloHandler controller.HelloHandler, // endpoint
	loginHandler controller.LoginHandler, // endpoint
	signupHandler controller.SignupController, // endpoint
) Router {
	return &router{
		secret:        string(secret),
		authUser:      authUser,
		helloHandler:  helloHandler,
		loginHandler:  loginHandler,
		signupHandler: signupHandler,
	}
}

func (r *router) RegisterEndpoint(e *echo.Echo) {
	// routes
	// free access
	e.GET("/", r.helloHandler.GetHello)
	e.POST("/signup", r.signupHandler.PostSignup)
	e.POST("/login", r.loginHandler.PostLogin)

	// restricted access(role user)
	user := e.Group("/user")
	user.Use(echojwt.WithConfig(echojwt.Config{SigningKey: []byte(r.secret)})) // jwt valid
	user.Use(r.authUser.AuthUser())                                            // user_id validation

	// restricted access(role adminonly)
	admin := e.Group("/admin")
	admin.Use(echojwt.WithConfig(echojwt.Config{SigningKey: []byte(r.secret)})) // jwt valid
	admin.Use(r.authUser.AuthAdminUser())                                       // user_id validation

}
