package usecase

import (
	"errors"

	"github.com/tetzng/golang-blog/model"
	"github.com/tetzng/golang-blog/repository"
	"github.com/tetzng/golang-blog/validator"
)

type UserUsecase interface {
	Login(user model.User) (*model.LoginUserResponse, error)
}

type userUsecase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

func (uu *userUsecase) Login(user model.User) (*model.LoginUserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return nil, err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return nil, err
	}
	if storedUser.Password != user.Password {
		return nil, errors.New("invalid password")
	}
	return &model.LoginUserResponse{
		Id:    storedUser.Id,
		Name:  storedUser.Name,
		Email: storedUser.Email,
	}, nil
}
