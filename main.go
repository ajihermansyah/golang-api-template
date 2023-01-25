package main

import (
	"golang-api-template/config/boot"
	cfg "golang-api-template/config/env"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

// App ...
type App struct {
	config cfg.Config
}

var app App

func init() {
	config := cfg.NewViperConfig()
	app = App{config: config}

	if config.GetBool(`app.debug`) {
		log.Println("Service RUN on DEBUG mode - HOST: " + config.GetString("app.host"))
	}
}

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	registerValidator()

	e := echo.New()
	apiHandler := boot.HTTPHandler{
		E:               e,
		Config:          app.config,
		ValidatorDriver: validatorDriver,
		Translator:      translator,
	}

	apiHandler.RegisterMiddleware()
	apiHandler.RegisterAPIHandler()

	e.Start(app.config.GetString(`app.host`))
}

func registerValidator() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	translator = trans

	validatorDriver = validator.New()
	enTranslations.RegisterDefaultTranslations(validatorDriver, translator)
}
