package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tetzng/golang-blog/model"
	"github.com/tetzng/golang-blog/usecase"
)

type UserController interface {
	Login(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu: uu}
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
