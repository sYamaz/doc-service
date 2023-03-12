package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	userDeleteOutputPort struct {
		c echo.Context
	}
)

func (h *userHandler) Delete(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	return h.s.DeleteUser(userId, &userDeleteOutputPort{c: c})
}

func (o *userDeleteOutputPort) Success() error {
	return o.c.NoContent(http.StatusOK)
}

func (o *userDeleteOutputPort) NotFound(err error) error {
	return notFoundError(o.c, err)
}
func (o *userDeleteOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}
