package controller

import (
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tetzng/golang-blog/model"
	"github.com/tetzng/golang-blog/usecase"
)

type TopController interface {
	Hello(c echo.Context) error
}

type topController struct {
	tu usecase.TopUsecase
}

func NewTopController(tu usecase.TopUsecase) TopController {
	return &topController{tu: tu}
}

func (tc *topController) Hello(c echo.Context) error {
	if t, ok := c.Get("user").(*jwt.Token); ok && t.Valid {
		user := model.User{}
		u, err := tc.tu.Hello(user, t)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "Hello, "+u.Name)
	}

	return c.String(http.StatusOK, "Hello, World")
}
