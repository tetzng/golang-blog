package router

import (
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetzng/golang-blog/controller"
)

func NewRouter(uc controller.UserController, tc controller.TopController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/login", uc.Login)
	e.POST("/sign_up", uc.SignUp)

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Skipper: func(c echo.Context) bool {
			ha := c.Request().Header.Get(echo.HeaderAuthorization)
			return strings.TrimSpace(ha) == ""
		},
	}))

	e.GET("/", tc.Hello)
	return e
}
