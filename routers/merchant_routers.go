package routers

import (
	"api-merchant-backend/config"
	"github.com/labstack/echo/v4"
	controller "api-merchant-backend/controllers/admin"
	service "api-merchant-backend/service/admin"
	repository "api-merchant-backend/repository/admin"
)

func MerchantRoutes(echo *echo.Echo, conf config.Config) {
	database := config.ConnectDB(conf)
	repo := repository.NewMerchantRepository(database)
	service := service.NewMerchantService(repo, conf)
	controllers := controller.Controller{
		Service: service,
	}

	merchant := echo.Group("/v1/admin")

	merchant.GET("/merchant", controllers.GetAllMerchant)
	merchant.GET("/merchant/:id", controllers.GetMerchantByID)
	merchant.POST("/merchant", controllers.InsertMerchant)
	merchant.PUT("/merchant/:id", controllers.UpdateMerchant)
	merchant.DELETE("/merchant/:id", controllers.DeleteMerchant)
}