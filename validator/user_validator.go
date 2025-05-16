package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/tetzng/golang-blog/model"
)

type UserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() UserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Email is not valid"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 100).Error("Password must be between 8 and 100 characters long"),
		),
	)
}
