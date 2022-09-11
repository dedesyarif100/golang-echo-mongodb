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
		return handleErrorValidator(err, context)
	}

	err = Controller.Service.Register(&admin)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     http.StatusBadRequest,
			"messages": err.Error(),
		})
	}
	return context.JSON(http.StatusCreated, map[string]interface{}{
		"code"		: http.StatusCreated,
		"messages"	: "success register admin",
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
		return handleErrorValidator(err, context)
	}

	result, er := Controller.Service.Login(&admin)
	if er != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"code"		: http.StatusBadRequest,
			"messages"	: er.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"code"		: http.StatusOK,
		"messages"	: "success login admin",
		"result"	: result,
	})
}

func handleErrorValidator(err error, context echo.Context) error {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
