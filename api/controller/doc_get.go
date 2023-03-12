package controller

import (
	"doc-api/api/usecase"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	docGetOutputPort struct {
		c echo.Context
	}
)

func (h *docHandler) Get(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	docId := c.Param("id")
	if docId == "" {
		return badrequestError(c, errors.New("path params is empty"))
	}

	return h.s.GetDoc(userId, docId, &docGetOutputPort{c: c})
}

func (o *docGetOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}

func (o *docGetOutputPort) Success(doc *usecase.Doc) error {
	type Item struct {
		Id    string `json:"document_id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	return o.c.JSON(http.StatusOK, Item{
		Id:    doc.Id,
		Title: doc.Title,
		Body:  doc.Body,
	})
}
