package http

import (
	"api-gorm-setting/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Health returns health check for the service.
func Health(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, "It is work v2!", nil)
	return echoCtx.JSON(res.Status, res)
}
