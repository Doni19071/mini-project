package controller

import (
	"net/http"
	"project/middleware"
	"project/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get All User
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
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

	if err := DB.Save(&seat).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new seat",
		"user":    seat,
	})
}

// Creat New Admin Account
func CreateAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new account",
		"user":    admin,
	})
}

// Login Admin
func LoginAdminController(c echo.Context) error {
	loginAdmin := models.Admin{}
	c.Bind(&loginAdmin)

	if err := DB.Where("username = ? AND password = ?", loginAdmin.Username, loginAdmin.Password).First(&loginAdmin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(loginAdmin.Username, loginAdmin.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{loginAdmin.Model, loginAdmin.Name, loginAdmin.Username, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}
