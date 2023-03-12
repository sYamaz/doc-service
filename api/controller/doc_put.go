package controller

import (
	"doc-api/api/usecase"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	docPutPutputPort struct {
		c echo.Context
	}
)

func (h *docHandler) Put(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	docId := c.Param("id")
	if docId == "" {
		return badrequestError(c, errors.New("path params is empty"))
	}

	type Body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	body := Body{}
	if err := c.Bind(&body); err != nil {
		return badrequestError(c, err)
	}
	return h.s.UpdateDoc(userId, usecase.Doc{
		Id:    docId,
		Title: body.Title,
		Body:  body.Body,
	}, &docPutPutputPort{c: c})
}

func (o *docPutPutputPort) Success() error {
	return o.c.NoContent(http.StatusOK)
}

func (o *docPutPutputPort) Failure(err error) error {
	return serverError(o.c, err)
}
