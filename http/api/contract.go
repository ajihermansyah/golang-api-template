package api

import (
	"golang-api-template/config/env"
	httpHelper "golang-api-template/helper"
	"golang-api-template/repository"
)

// InjectAPIHandler ...
type InjectAPIHandler struct {
	Config   env.Config
	Helper   httpHelper.HTTPHelper
	UserRepo repository.UserRepositoryInterface
}
