package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	signupController struct {
		service usecase.SignupService
	}

	signupOutputPort struct{ ctx echo.Context }

	SignupController interface {
		PostSignup(ctx echo.Context) error
	}
)

func NewSignupHandler(service usecase.SignupService) SignupController {
	return &signupController{
		service: service,
	}
}

func (c *signupController) PostSignup(ctx echo.Context) error {
	type Body struct {
		userId   string
		password string
	}

	body := new(Body)
	if err := ctx.Bind(body); err != nil {
		return err
	}

	return c.service.Signup(body.userId, body.password, &signupOutputPort{ctx: ctx})
}

func (o *signupOutputPort) Success() error {
	return o.ctx.NoContent(http.StatusOK)
}

func (o *signupOutputPort) Failure(err error) error {
	return serverError(o.ctx, err)
}
