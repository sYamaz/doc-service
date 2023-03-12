package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	userPutOutputPort struct{ c echo.Context }
)

func (h userHandler) Put(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	type Body struct {
		UserId   string
		Password string
	}

	body := new(Body)
	if err := c.Bind(body); err != nil {
		return badrequestError(c, err)
	}
	if body.UserId != userId {
		return badrequestError(c, err)
	}

	return h.s.UpdateUser(&usecase.UserInfo{
		Id:       body.UserId,
		Password: body.Password,
	}, &userPutOutputPort{c: c})
}

func (o userPutOutputPort) Success() error {
	return o.c.NoContent(http.StatusOK)
}

func (o userPutOutputPort) NotFound(err error) error {
	return notFoundError(o.c, err)
}
func (o userPutOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}
