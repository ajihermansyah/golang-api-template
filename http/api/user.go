package api

import (
	"golang-api-template/config/env"
	"golang-api-template/helper"
	httpHelper "golang-api-template/helper"
	defaultValue "golang-api-template/helper/defaultvalue"
	"golang-api-template/model/entity"
	"golang-api-template/model/request"
	resp "golang-api-template/model/response"
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

// create user handler
func (_h *UserApiHandler) CreateUserHandler(c echo.Context) error {
	var (
		err     error
		input   request.UserRequest
		userObj entity.User
	)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	if err = c.Bind(&input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	//  for validate struct request input
	if err := validator.New().Struct(input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// for check input custom validation
	if err := input.CustomValidate(); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// set mapping value input to entity objects
	err = helper.Automapper(input, &userObj)
	if err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// set value objs needed
	userObj.CreatedAt = now

	// cleaning data objs
	userObj = userObj.DataCleaning()

	err = _h.UserRepo.CreateUser(userObj)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "failed to create user", err)
	}

	return _h.Helper.SendSuccess(c, "Success", _h.Helper.EmptyJsonMap())
}

// get user handler
func (_h *UserApiHandler) GetUserHandler(c echo.Context) error {
	var (
		err         error
		limit, page int
	)

	// atribute for search
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	keyword := c.QueryParam("keywords")
	filterText := c.QueryParam("filter_text")

	limit, page = defaultValue.SetDafaultValuePagination(limitStr, pageStr)

	// get data user
	users, totalRecord, err := _h.UserRepo.FindAllUser(limit, page, filterText, keyword)
	if err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// add data to pagging
	pagination := _h.Helper.GeneratePaging(c, 0, 0, limit, page, totalRecord)

	// // don't use return response using json map string interface
	// return c.JSON(200, map[string]interface{}{
	// 	"code":         200,
	// 	"code_type":    "success",
	// 	"code_message": "Success",
	// 	"data":         result,
	// 	"pagination":   pagination,
	// })

	responseObj := resp.APIWithPaginationResponse{
		Data:       users,
		Pagination: pagination,
	}

	return _h.Helper.SendSuccess(c, "Success", responseObj)
}

// update user handler
func (_h *UserApiHandler) UpdateUserHandler(c echo.Context) error {
	var err error
	var input request.UserRequest

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	userId := c.Param("user_id")

	if err = c.Bind(&input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	//  for validate struct request input
	if err := validator.New().Struct(input); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// for check input custom validation
	if err := input.CustomValidate(); err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	user, _ := _h.UserRepo.FindUserById(userId)
	if user.ID == "" {
		return _h.Helper.SendNotFoundError(c, "User not found", _h.Helper.EmptyJsonMap())
	}

	// set mapping value input to entity objects
	err = helper.Automapper(input, &user)
	if err != nil {
		return _h.Helper.SendBadRequest(c, err.Error(), _h.Helper.EmptyJsonMap())
	}

	// set value objs needed
	user.UpdatedAt = now

	// cleaning data objs
	user = user.DataCleaning()

	err = _h.UserRepo.UpdateUser(user)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "failed to update user", err)
	}

	return _h.Helper.SendSuccess(c, "Success", _h.Helper.EmptyJsonMap())
}

// get detail user handler
func (_h *UserApiHandler) DetailUserHandler(c echo.Context) error {
	userId := c.Param("user_id")

	user, _ := _h.UserRepo.FindUserById(userId)
	if user.ID == "" {
		return _h.Helper.SendNotFoundError(c, "User not found", _h.Helper.EmptyJsonMap())
	}

	return _h.Helper.SendSuccess(c, "Success", user)
}

// delete user handler
func (_h *UserApiHandler) DeleteUserHandler(c echo.Context) error {
	var err error

	userId := c.Param("user_id")

	// check user data
	user, _ := _h.UserRepo.FindUserById(userId)
	if user.ID == "" {
		return _h.Helper.SendNotFoundError(c, "User not found", _h.Helper.EmptyJsonMap())
	}

	err = _h.UserRepo.DeleteUserByID(userId)
	if err != nil {
		return _h.Helper.SendBadRequest(c, "failed to delete user", err)
	}

	return _h.Helper.SendSuccess(c, "Success", _h.Helper.EmptyJsonMap())
}
