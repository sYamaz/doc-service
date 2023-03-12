package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	userGetOutputPort struct{ c echo.Context }
)

func (h *userHandler) Get(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}
	return h.s.GetUser(userId, &userGetOutputPort{c: c})
}

func (o *userGetOutputPort) Success(info *usecase.UserInfo) error {
	type User struct {
		Id       string
		Password string
	}

	return o.c.JSON(http.StatusOK, User{
		Id:       "id",
		Password: "****************",
	})
}

func (o *userGetOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}

func (o *userGetOutputPort) NotFound(err error) error {
	return notFoundError(o.c, err)
}
