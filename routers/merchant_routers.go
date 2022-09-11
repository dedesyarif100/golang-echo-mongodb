package routers

import (
	"api-merchant-backend/config"
	controller "api-merchant-backend/controllers/admin"
	repository "api-merchant-backend/repository/admin"
	service "api-merchant-backend/service/admin"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MerchantRoutes(echo *echo.Echo, conf config.Config) {
	database := config.ConnectDB(conf)
	repo := repository.NewMerchantRepository(database)
	service := service.NewMerchantService(repo, conf)
	controllers := controller.Controller{
		Service: service,
	}

	secret := os.Getenv("JWT_SECRET")

	authMerchant := echo.Group("/v1/admin")
	authMerchant.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey		: []byte(secret),
		SigningMethod	: "HS256",
	}))

	authMerchant.GET("/merchant", controllers.GetAllMerchant)
	authMerchant.GET("/merchant/:id", controllers.GetMerchantByID)
	authMerchant.POST("/merchant", controllers.InsertMerchant)
	authMerchant.PUT("/merchant/:id", controllers.UpdateMerchant)
	authMerchant.DELETE("/merchant/:id", controllers.DeleteMerchant)
}