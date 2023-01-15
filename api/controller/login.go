package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	loginHandler struct {
		service usecase.LoginService
	}
	LoginHandler interface {
		PostLogin(ctx echo.Context) error
	}

	loginOutputPort struct {
		ctx echo.Context
	}
)

func NewLoginHandler(service usecase.LoginService) LoginHandler {
	return &loginHandler{
		service: service,
	}
}

func (h *loginHandler) PostLogin(ctx echo.Context) error {
	type Body struct {
		userId   string
		password string
	}

	// body paramの抽出
	body := new(Body)
	if err := ctx.Bind(body); err != nil {
		return err
	}

	return h.service.Login(body.userId, body.password, &loginOutputPort{ctx: ctx})
}

func (o *loginOutputPort) Success(token string, admin bool, err error) error {
	type Body struct {
		Token string
		Admin bool
	}
	return o.ctx.JSON(http.StatusOK, Body{Token: token, Admin: admin})
}

func (o *loginOutputPort) InfoMissmatched(err error) error {
	return authError(o.ctx, err)
}

func (o *loginOutputPort) Failure(err error) error {
	return authError(o.ctx, err)
}
