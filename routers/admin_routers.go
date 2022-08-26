package routers

import (
	"api-merchant-backend/config"
	"github.com/labstack/echo/v4"
	controller "api-merchant-backend/controllers/auth"
	service "api-merchant-backend/service/auth"
	repository "api-merchant-backend/repository/auth"
)

func AdminRouters(echo *echo.Echo, conf config.Config) {
	database := config.ConnectDB(conf)
	repo := repository.NewAdminRepository(database)
	service := service.NewAdminService(repo, conf)
	controllers := controller.Controller{
		Service: service,
	}

	admin := echo.Group("/v1/admin")
	admin.POST("/register", controllers.Register)
	admin.POST("/login", controllers.Login)
}