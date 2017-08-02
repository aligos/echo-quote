package quoteroute

import (
	"github.com/aligos/echo-quote/api/quote/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/api/quotes", quotecontroller.GetAll)
	e.GET("/api/quotes/:id", quotecontroller.GetById)
	e.POST("/api/quotes", quotecontroller.NewQuote)
	e.DELETE("/api/quotes/:id", quotecontroller.RemoveQuote)
}
