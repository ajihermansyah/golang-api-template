package api

import (
	"github.com/labstack/echo/v4"
)

// PingHandler ...
func (_h *InjectAPIHandler) PingHandler(c echo.Context) error {
	return _h.Helper.SendSuccess(c, "EVERYTHING IS WORKING FINE...", _h.Helper.EmptyJsonMap())
}
