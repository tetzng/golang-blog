package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetzng/golang-blog/controller"
)

func NewRouter(uc controller.UserController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/login", uc.Login)
	return e
}
