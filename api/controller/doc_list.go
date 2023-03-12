package controller

import (
	"doc-api/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	docListOutputPort struct {
		c echo.Context
	}
)

func (h *docHandler) List(c echo.Context) error {
	userId, err := extractUserId(c)
	if err != nil {
		return authError(c, err)
	}

	return h.s.GetDocs(userId, &usecase.DocListCondition{OnlyOwn: false}, &docListOutputPort{c: c})
}

func (o *docListOutputPort) Failure(err error) error {
	return serverError(o.c, err)
}

func (o *docListOutputPort) Success(docs []usecase.DocHeader) error {
	type Item struct {
		Id    string `json:"document_id"`
		Title string `json:"title"`
	}
	type ListBody struct {
		Docs []Item `json:"documents"`
	}

	body := ListBody{}
	for _, doc := range docs {
		body.Docs = append(body.Docs, Item{Id: doc.Id, Title: doc.Title})
	}
	return o.c.JSON(http.StatusOK, body)
}
