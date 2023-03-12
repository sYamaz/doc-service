package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	helloHandler struct{}

	HelloHandler interface {
		Get(ctx echo.Context) error
	}
)

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Get(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello world!")
}
