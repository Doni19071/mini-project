package appController

import (
	"fmt"
	"net/http"

	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	model     appModel.OrderModel
	jwtSecret string
}

func NewOrderController(jwtSecret string, m appModel.OrderModel) OrderController {
	return OrderController{
		jwtSecret: jwtSecret,
		model:     m,
	}
}

func (pc OrderController) AddOrder(c echo.Context) error {
	var order appModel.Order
	if err := c.Bind(&order); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid order data")
	}
	order, err := pc.model.AddOrder(order)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add order")
	}
	return c.JSON(http.StatusOK, order)
}

func (pc OrderController) GetAllOrder(c echo.Context) error {
	allOrder, err := pc.model.GetAllOrder()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get orders")
	}
	return c.JSON(http.StatusOK, allOrder)
}
