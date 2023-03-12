package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	docDeleteOutputPort struct {
		c echo.Context
	}
)

func (h *docHandler) Delete(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	docId := c.Param("id")
	if docId == "" {
		return badrequestError(c, errors.New("path params is empty"))
	}

	return h.s.DeleteDoc(userId, docId, &docDeleteOutputPort{c: c})
}

func (o *docDeleteOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}

func (o *docDeleteOutputPort) Success() error {
	return o.c.NoContent(http.StatusOK)
}
