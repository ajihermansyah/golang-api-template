package repository

import (
	"golang-api-template/model/entity"
)

type UserRepositoryInterface interface {
	CreateUser(user entity.User) error
	FindAllUser(page int, limit int, filterText, keyword string) ([]entity.User, int, error)
}
