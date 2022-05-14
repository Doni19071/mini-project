package routes

import (
	"project/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	// API User
	e.POST("/api/v1/login", controller.LoginUserController)
	e.POST("/api/v1/register", controller.CreateUserController)
	e.GET("/api/v1/seats", controller.GetSeatsController)

	// API Admin
	e.POST("/api/admin/register", controller.CreateAdminController)
	e.POST("/api/admin/login", controller.LoginAdminController)
	e.GET("/api/v1/users", controller.GetUsersController)
	e.POST("/api/v1/seats", controller.CreateSeatController)
	e.DELETE("/api/v1/seats", controller.DeleteSeatController)

	return e
}
