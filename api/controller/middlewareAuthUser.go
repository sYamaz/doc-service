package controller

import (
	"doc-api/api/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type (
	AuthUserMiddleware interface {
		AuthUser() echo.MiddlewareFunc
		AuthAdminUser() echo.MiddlewareFunc
	}
	authUserMiddleware struct {
		t usecase.AuthUser
	}

	authUserOutputPort struct {
		next echo.HandlerFunc
		ctx  echo.Context
	}
)

func NewAuthUserMiddleware(t usecase.AuthUser) AuthUserMiddleware {
	return &authUserMiddleware{
		t: t,
	}
}

func (a *authUserMiddleware) AuthUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			return a.t.Auth(user, &authUserOutputPort{
				next: next,
				ctx:  c,
			})
		}
	}
}

func (a *authUserMiddleware) AuthAdminUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			return a.t.AuthAsAdmin(user, &authUserOutputPort{
				next: next,
				ctx:  c,
			})
		}
	}
}

func (o *authUserOutputPort) Success(userId string) error {
	o.ctx.Set("user_id", userId)
	return o.next(o.ctx)
}

func (o *authUserOutputPort) Failure(err error) error {
	return authError(o.ctx, err)
}
