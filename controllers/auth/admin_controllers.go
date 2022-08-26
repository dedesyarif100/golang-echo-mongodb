package auth

import (
	"net/http"
	"api-merchant-backend/entity"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/service/auth"
)

type Controller struct {
	Service auth.AdminService
}

func (Controller *Controller) Register(context echo.Context) error {
	var admin entity.AuthRegister
	err := context.Bind(&admin)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if err := context.Validate(admin); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	_, er := Controller.Service.Register(&admin)
	if er != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": er.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register account",
	})
}

func (Controller *Controller) Login(context echo.Context) error {
	var admin entity.AuthLogin
	err := context.Bind(&admin)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if err := context.Validate(admin); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	result, er := Controller.Service.Login(&admin)
	if er != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": er.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login account",
		"result":   result,
	})
}