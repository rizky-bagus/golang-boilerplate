package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(barangHandler *BarangHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.GET("/", Status)
	engine.GET("/healthz", Health)
	engine.GET("/version", Version)

	engine.POST("/create-barang", barangHandler.CreateBarang)
	engine.GET("/list-barang", barangHandler.GetListBarang)
	engine.GET("/get-barang/:id", barangHandler.GetDetailBarang)
	engine.PUT("/update-barang/:id", barangHandler.UpdateBarang)
	engine.DELETE("/delete-barang/:id", barangHandler.DeleteBarang)

	return engine
}
