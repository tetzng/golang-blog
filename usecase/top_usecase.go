package usecase

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/tetzng/golang-blog/model"
	"github.com/tetzng/golang-blog/repository"
)

type TopUsecase interface {
	Hello(user model.User, t *jwt.Token) (*model.UserResponse, error)
}
type topUsecase struct {
	ur repository.UserRepository
}

func NewTopUsecase(ur repository.UserRepository) TopUsecase {
	return &topUsecase{ur: ur}
}

func (tu *topUsecase) Hello(user model.User, t *jwt.Token) (*model.UserResponse, error) {
	claims := t.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("user_id not found in token")
	}

	floatUserId, ok := userId.(float64)
	if !ok {
		return nil, fmt.Errorf("invalid user_id in token")
	}

	if err := tu.ur.GetUserById(&user, uint(floatUserId)); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
