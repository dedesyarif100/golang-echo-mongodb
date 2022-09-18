package scores

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/service/scores"
)

type Controller struct {
	Service scores.ScoresService
}

func (controller *Controller) GetAllScores(context echo.Context) error {
	results, num, err := controller.Service.GetAllScores()
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"message"	: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"code"		: http.StatusOK,
		"message"	: "success get all scores",
		"result"	: results,
		"result_number": num,
	})
}
