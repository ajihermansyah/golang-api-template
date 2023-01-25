package repository

import (
	"golang-api-template/model/entity"
)

type UserRepositoryInterface interface {
	CreateUser(user entity.User) error
}
