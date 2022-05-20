package appController

import (
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo, jwtSecret string, adminModel appModel.AdminModel, seatModel appModel.SeatModel, orderModel appModel.OrderModel) (AdminController, SeatController, OrderController) {
	adminController := NewAdminController(jwtSecret, adminModel)
	seatController := NewSeatController(jwtSecret, seatModel)
	orderController := NewOrderController(jwtSecret, orderModel)

	e.POST("/api/v1/register", adminController.AddAdmin)
	e.POST("/api/v1/login", adminController.Login)
	e.GET("/api/v1/seats", seatController.GetAllSeat)
	e.POST("/api/v1/orders", orderController.AddOrder)

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))

	e.GET("/api/v1/admin/orders", orderController.GetAllOrder, jwtMiddleware)
	e.POST("/api/v1/admin/seats", seatController.AddSeat, jwtMiddleware)

	return adminController, seatController, orderController
}
