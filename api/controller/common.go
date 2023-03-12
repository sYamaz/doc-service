package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 401
func authError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusUnauthorized, err)
}

// 400
func badrequestError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, err)
}

// 500
func serverError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

// 404
func notFoundError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusNotFound, err)
}

// util
func extractUserId(c echo.Context) (string, error) {
	id, ok := c.Get("user_id").(string)
	if !ok {
		return "", errors.New("user_id is not exists in Context")
	}

	return id, nil
}
