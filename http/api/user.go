package api

import (
	"encoding/json"
	"fmt"
	"golang-api-template/config/env"
	httpHelper "golang-api-template/helper"
	"golang-api-template/model/entity"
	"golang-api-template/model/request"
	"golang-api-template/repository"
	"time"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/v4"
)

type UserApiHandler struct {
	Helper   httpHelper.HTTPHelper
	Config   env.Config
	UserRepo repository.UserRepositoryInterface
}

// create user
func (_h *UserApiHandler) CreateUser(c echo.Context) error {
	var (
		err   error
		input request.UserRequest
	)

	if err = c.Bind(&input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	if err := validator.New().Struct(input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("CREATE USER JSON REQUEST PAYLOAD = ", string(b))

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	user := entity.User{
		Name:      input.Name,
		Email:     input.Email,
		Gender:    input.Gender,
		Age:       input.Age,
		CreatedAt: now,
	}

	err = _h.UserRepo.CreateUser(user)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "failed to create user", err)
	}

	return _h.Helper.SendSuccess(c, "Success", _h.Helper.EmptyJsonMap())
}
