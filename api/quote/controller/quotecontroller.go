package quotecontroller

import (
	"net/http"

	"github.com/aligos/echo-quote/api/quote/actions"
	quote "github.com/aligos/echo-quote/api/quote/model"
	"github.com/labstack/echo"
)

func GetRandomQuote(c echo.Context) error {
	q, _ := quoteactions.GetRandomQuote()

	return c.JSON(http.StatusOK, q)
}

func GetAll(c echo.Context) error {
	qs, _ := quoteactions.All()

	return c.JSON(http.StatusOK, qs)
}

func GetById(c echo.Context) error {
	id := c.Param("id")

	nq, _ := quoteactions.GetById(id)

	return c.JSON(http.StatusOK, nq)
}

func NewQuote(c echo.Context) error {
	q := new(quote.Quote)

	c.Bind(q)

	nq, _ := quoteactions.NewQuote(*q)

	return c.JSON(http.StatusOK, nq)
}

func RemoveQuote(c echo.Context) error {
	id := c.Param("id")

	quoteactions.DeleteQuote(id)

	return c.String(http.StatusOK, "")
}
