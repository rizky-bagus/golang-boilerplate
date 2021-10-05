package http

import (
	"api-gorm-setting/entity"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Version returns health check for the service.
func Version(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, os.Getenv("APP_VERSION"), nil)
	return echoCtx.JSON(res.Status, res)
}
