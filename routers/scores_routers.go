package routers

import (
	"os"
	"github.com/labstack/echo/v4"
	"api-merchant-backend/config"
	controller "api-merchant-backend/controllers/scores"
	repository "api-merchant-backend/repository/scores"
	service "api-merchant-backend/service/scores"
	"github.com/labstack/echo/v4/middleware"
)

func ScoresRouters(echo *echo.Echo, conf config.Config) {
	database := config.ConnectDB(conf)
	repo := repository.NewScoresRepository(database)
	service := service.NewScoresService(repo, conf)
	controllers := controller.Controller{
		Service: service,
	}

	secret := os.Getenv("JWT_SECRET")

	scores := echo.Group("/v1/scores")
	scores.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey		: []byte(secret),
		SigningMethod	: "HS256",
	}))
	scores.GET("/scores", controllers.GetAllScores)
}