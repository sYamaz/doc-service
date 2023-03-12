package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	docPostOutputPort struct {
		c echo.Context
	}
)

func (h *docHandler) Post(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	type Body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	body := Body{}
	if err = c.Bind(&body); err != nil {
		return badrequestError(c, err)
	}

	return h.s.PostDoc(userId, usecase.Doc{Title: body.Title, Body: body.Body}, &docPostOutputPort{c: c})
}

func (o *docPostOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}

func (o *docPostOutputPort) Success(docId string) error {
	return o.c.JSON(http.StatusCreated, struct {
		DocId string `json:"document_id"`
	}{
		DocId: docId,
	})
}
