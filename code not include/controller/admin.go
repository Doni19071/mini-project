package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/loginJWT"
	"project/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get All Order
func GetOrdersController(c echo.Context) error {
	var orders []models.Order

	if err := config.DB.Find(&orders).Error; err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all orders",
		"orders":  orders,
	})
}

// Delete Seat by ID
func DeleteSeatController(c echo.Context) error {
	var seats = map[int]*models.Seat{}
	id, _ := strconv.Atoi(c.Param("id"))
	delete(seats, id)
	return c.NoContent(http.StatusNoContent)
}

// Creat New Seat
func CreateSeatController(c echo.Context) error {
	seat := models.Seat{}
	c.Bind(&seat)

	if err := config.DB.Save(&seat).Error; err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new seat",
		"seat":    seat,
	})
}

// Creat New Admin Account
func CreateAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new account",
		"admin":   admin,
	})
}

// Login Admin
func LoginAdminController(c echo.Context) error {
	loginAdmin := models.Admin{}
	c.Bind(&loginAdmin)

	if err := config.DB.Where("username = ? AND password = ?", loginAdmin.Username, loginAdmin.Password).First(&loginAdmin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := loginJWT.CreateJwtToken(loginAdmin.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}
