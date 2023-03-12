package controller

import (
	"doc-api/api/usecase"

	"github.com/labstack/echo/v4"
)

type (
	DocHandler interface {
		List(c echo.Context) error
		Get(c echo.Context) error
		Post(c echo.Context) error
		Put(c echo.Context) error
		Delete(c echo.Context) error
	}
	docHandler struct {
		s usecase.DocService
	}
)

func NewDocHandler(s usecase.DocService) DocHandler {
	return &docHandler{s: s}
}
