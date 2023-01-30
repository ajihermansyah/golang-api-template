package request

import (
	errorHelper "golang-api-template/helper/error"
	"strings"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	Age      int    `json:"age" validate:"required,numeric"`
	Email    string `json:"email" validate:"required,email"`
	Gender   string `json:"gender" validate:"required,alpha"`
}

func (user UserRequest) CustomValidate() error {

	if len(strings.TrimSpace(user.Name)) == 0 {
		return errorHelper.FullnameMustNotEmpty
	}
	if len(strings.TrimSpace(user.Username)) == 0 {
		return errorHelper.UsernameMustNotEmpty
	}
	if len(strings.TrimSpace(user.Password)) == 0 {
		return errorHelper.PasswordMustNotEmpty
	}
	if user.Age == 0 {
		return errorHelper.AgeMustNotEmpty
	}
	if len(strings.TrimSpace(user.Gender)) == 0 {
		return errorHelper.GenderMustNotEmpty
	}
	if len(strings.TrimSpace(user.Email)) == 0 {
		return errorHelper.EmailMustNotEmpty
	}

	return nil
}
