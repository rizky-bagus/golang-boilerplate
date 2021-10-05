package http

import (
	"api-gorm-setting/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Status returns health check for the service.
func Status(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, "It is work v2!", nil)
	return echoCtx.JSON(res.Status, res)
}
