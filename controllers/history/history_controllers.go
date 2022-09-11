package history

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/service/history"
)

type Controller struct {
	Service history.HistoryService
}

func (controller *Controller) GetCountPointByIdCoupon(context echo.Context) error {
	results, err := controller.Service.GetCountPointByIdCoupon()
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"message"	: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"code"		: http.StatusOK,
		"message"	: "success get all merchant",
		"result"	: results,
	})
}