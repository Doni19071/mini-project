package appController

import (
	"fmt"
	"net/http"

	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

type SeatController struct {
	model     appModel.SeatModel
	jwtSecret string
}

func NewSeatController(jwtSecret string, m appModel.SeatModel) SeatController {
	return SeatController{
		jwtSecret: jwtSecret,
		model:     m,
	}
}

func (pc SeatController) AddSeat(c echo.Context) error {
	var seat appModel.Seat
	if err := c.Bind(&seat); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid seat data")
	}
	seat, err := pc.model.AddSeat(seat)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add seat")
	}
	return c.JSON(http.StatusOK, seat)
}

func (pc SeatController) GetAllSeat(c echo.Context) error {
	allSeat, err := pc.model.GetAllSeat()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get seats")
	}
	return c.JSON(http.StatusOK, allSeat)
}
