package appController

import (
	"fmt"
	"net/http"
	"strconv"

	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminController struct {
	model     appModel.AdminModel
	jwtSecret string
}

func NewAdminController(jwtSecret string, m appModel.AdminModel) AdminController {
	return AdminController{
		jwtSecret: jwtSecret,
		model:     m,
	}
}

func (pc AdminController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	admin, err := pc.model.GetUsernameAndPassword(loginInfo.Username, loginInfo.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}
	token, err := appMiddleware.CreateToken(int(admin.ID), pc.jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}
	admin.Token = token
	admin, err = pc.model.EditAdmin(int(admin.ID), admin)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add token")
	}
	return c.JSON(http.StatusOK, admin)
}

func (pc AdminController) AddAdmin(c echo.Context) error {
	var admin appModel.Admin
	if err := c.Bind(&admin); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid admin data")
	}
	admin, err := pc.model.AddAdmin(admin)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add admin")
	}
	return c.JSON(http.StatusOK, admin)
}

func (pc AdminController) EditAdmin(c echo.Context) error {
	adminId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid admin id")
	}
	var admin appModel.Admin
	if err := c.Bind(&admin); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid admin data")
	}
	admin, err = pc.model.EditAdmin(adminId, admin)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit admin")
	}
	return c.JSON(http.StatusOK, admin)
}

//coba edit
