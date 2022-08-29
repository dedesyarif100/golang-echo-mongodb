package auth

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/entity"
	"api-merchant-backend/service/auth"
	"github.com/go-playground/validator"
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
		return handleError(err, context)
	}

	result, er := Controller.Service.Register(&admin)
	if er != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": er.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register account",
		"result":   result,
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
		return handleError(err, context)
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

func handleError(err error, context echo.Context) error {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "min":
					report.Message = fmt.Sprintf("%s min password 8 character", err.Field())
			}
			break
		}
	}

	return context.JSON(report.Code, report)
}
