package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ErrorBody struct {
		Message string
	}
)

func authError(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, ErrorBody{Message: err.Error()})
}

func serverError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, ErrorBody{Message: err.Error()})
}
