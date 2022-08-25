package admin

import (
	"net/http"
	"api-merchant-backend/entity"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/service/admin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	Service admin.MerchantService
}

func (controller *Controller) InsertMerchant(context echo.Context) error {
	merchant := entity.Merchant{
		ID: primitive.NewObjectID(),
		Information: entity.Information{},
	}

	err := context.Bind(&merchant)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	merchant, er := controller.Service.InsertMerchant(merchant)
	if er != nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"messages": er.Error(),
		})
	}

	return context.JSON(http.StatusCreated, map[string]any{
		"messages": "successfully created reward card coupon",
		"data":     merchant,
	})
}


func (controller *Controller) GetAllMerchant(context echo.Context) error {
	results, err := controller.Service.GetAllMerchant()
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"messages"	: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"code"		: http.StatusOK,
		"message"	: "success get all merchant",
		"result"	: results,
	})
}


func (controller *Controller) GetMerchantByID(context echo.Context) error {
	id := context.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := controller.Service.GetMerchantByID(objID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"message"	: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"code"		: http.StatusOK,
		"message"	: "success get merchant by id",
		"result"	: result,
	})
}

func (controller *Controller) UpdateMerchant(context echo.Context) error {
	id := context.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	data := entity.Merchant{
		ID: objID,
	}
	err := context.Bind(&data)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"message"	: err.Error(),
		})
	}

	result, err := controller.Service.UpdateMerchant(objID, &data)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]any{
			"code"		: http.StatusBadRequest,
			"message"	: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"code"		: http.StatusOK,
		"message"	: "success update merchant",
		"result"	: result.ID,
	})
}

func (controller *Controller) DeleteMerchant(context echo.Context) error {
	id := context.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	err := controller.Service.DeleteMerchant(objID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, map[string]any{
		"message": "success delete merchant",
	})
}
