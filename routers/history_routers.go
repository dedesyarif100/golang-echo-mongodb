package routers

import (
	"api-merchant-backend/config"
	"github.com/labstack/echo/v4"
	controller "api-merchant-backend/controllers/history"
	repository "api-merchant-backend/repository/history"
	service "api-merchant-backend/service/history"
)

func HistoryRoutes(echo *echo.Echo, conf config.Config) {
	database := config.ConnectDB(conf)
	repo := repository.NewHistoryRepository(database)
	service := service.NewHistoryService(repo, conf)
	controllers := controller.Controller{
		Service: service,
	}

	authHistory := echo.Group("/v1/admin")
	authHistory.GET("/countpoint", controllers.GetCountPointByIdCoupon)
}