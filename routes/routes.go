package routes

import (
	"github.com/aligos/echo-quote/api/quote/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	quoteroute.Init(e)
}
