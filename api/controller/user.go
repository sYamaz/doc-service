package controller

import (
	"doc-api/api/usecase"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
		Post(c echo.Context) error
		Get(c echo.Context) error
		Put(c echo.Context) error
		Delete(c echo.Context) error
	}
	userHandler struct{ s usecase.UserService }
)

func NewUserHandler(s usecase.UserService) UserHandler {
	return &userHandler{s: s}
}
