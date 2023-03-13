package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	userPostOutputPort struct {
		c echo.Context
	}
)

func (h *userHandler) Post(c echo.Context) error {
	type Body struct {
		UserId   string `json:"user_id"`
		Password string `json:"password"`
	}

	body := new(Body)
	if err := c.Bind(body); err != nil {
		return badrequestError(c, err)
	}

	return h.s.SignupUser(&usecase.UserInfo{
		Id:       body.UserId,
		Password: body.Password,
	}, &userPostOutputPort{c: c})
}

func (o *userPostOutputPort) Success() error {
	return o.c.NoContent(http.StatusOK)
}

func (o *userPostOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}
