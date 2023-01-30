package boot

import (
	"golang-api-template/config/env"
	httpHelper "golang-api-template/helper"
	"golang-api-template/http/api"
	"golang-api-template/model/mongo"

	"golang-api-template/repository/user"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// HTTPHandler ...
type HTTPHandler struct {
	E               *echo.Echo
	Config          env.Config
	Helper          httpHelper.HTTPHelper
	ValidatorDriver *validator.Validate
	Translator      ut.Translator
}

// RegisterAPIHandler ...
func (h *HTTPHandler) RegisterAPIHandler() *HTTPHandler {
	h.Helper = httpHelper.HTTPHelper{
		Validate:   h.ValidatorDriver,
		Translator: h.Translator,
	}

	//API routes handler
	dbMongo := mongo.Info{
		Hostname: h.Config.GetString("database.mongodb.host"),
		Database: h.Config.GetString("database.mongodb.database"),
		Username: h.Config.GetString("database.mongodb.username"),
		Password: h.Config.GetString("database.mongodb.password"),
	}

	// model initialize
	dbConnect, err := dbMongo.Connect()
	if err != nil {
		panic(err.Error())
	}

	userRepo := user.NewUserRepository(dbConnect, dbMongo.Database)
	apiHandler := api.InjectAPIHandler{
		Config: h.Config,
		Helper: h.Helper,
	}

	userHandler := &api.UserApiHandler{
		Helper:   h.Helper,
		Config:   h.Config,
		UserRepo: userRepo,
	}

	router := h.E
	group := router.Group(`api/v1`)

	//ping
	router.GET("/ping", apiHandler.PingHandler)

	group.POST("/users", userHandler.CreateUser)
	group.GET("/users", userHandler.GetUser)

	return h
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Golang-API-Template/1.0")
		return next(c)
	}
}

// RegisterMiddleware ...
func (h *HTTPHandler) RegisterMiddleware() {
	h.E.Use(serverHeader)
	h.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	h.E.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	if h.Config.GetBool(`app.debug`) == true {
		h.E.Use(middleware.Logger())
		h.E.HideBanner = true
		h.E.Debug = true
	} else {
		h.E.HideBanner = true
		h.E.Debug = false
		h.E.Use(middleware.Recover())
	}
}
