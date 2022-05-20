package routes

import (
	"project/controller"
	"project/loginJWT"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	// API Order
	e.GET("/api/v1/seats", controller.GetSeatsController)
	e.POST("/api/v1/orders", controller.CreateOrderController)

	// API Admin
	e.POST("/api/v1/admin/register", controller.CreateAdminController)
	e.POST("/api/v1/admin/login", controller.LoginAdminController)
	e.GET("/api/v1/admin/orders", controller.GetOrdersController, middleware.JWT([]byte(loginJWT.SECRET_JWT)))
	e.POST("/api/v1/admin/seats", controller.CreateSeatController, middleware.JWT([]byte(loginJWT.SECRET_JWT)))
	e.DELETE("/api/v1/admin/seats", controller.DeleteSeatController, middleware.JWT([]byte(loginJWT.SECRET_JWT)))

	return e
}
