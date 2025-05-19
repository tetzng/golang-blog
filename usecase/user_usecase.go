package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tetzng/golang-blog/model"
	"github.com/tetzng/golang-blog/repository"
	"github.com/tetzng/golang-blog/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Login(user model.User) (string, error)
	SignUp(user model.User) error
}

type userUsecase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.Id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	ts, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return ts, nil
}

func (uu *userUsecase) SignUp(user model.User) error {
	if err := uu.uv.UserValidate(user); err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return err
	}
	return nil
}
