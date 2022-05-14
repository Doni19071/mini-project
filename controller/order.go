package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/models"

	//	"strconv"

	"github.com/labstack/echo/v4"
)

// Get All Seat
func GetSeatsController(c echo.Context) error {
	var seats = []models.Seat{}

	if err := config.DB.Find(&seats).Error; err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all seats",
		"seats":   seats,
	})
}

// Creat New Order
func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	if err := config.DB.Save(&order).Error; err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"order":   order,
	})
}
