package controller

import (
	"net/http"
	"project/middleware"
	"project/models"

	//	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Get All Seat
func GetSeatsController(c echo.Context) error {
	var seats []models.Seat

	if err := DB.Find(&seats).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   seats,
	})
}

// Creat New Account
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new account",
		"user":    user,
	})
}

// Login User
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.Username, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{user.Model, user.Name, user.Username, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}

// // Delete Order Seat by ID
// func DeleteOrderSeatController(c echo.Context) error {
// 	var seats = map[int]*models.Order{}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(seats, id)
// 	return c.NoContent(http.StatusNoContent)
// }

// Creat New Account
func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	if err := DB.Save(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"user":    order,
	})
}
