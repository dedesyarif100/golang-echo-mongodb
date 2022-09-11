package routers

import (
	"api-merchant-backend/config"
	c "api-merchant-backend/controllers/reward"
	r "api-merchant-backend/repository/reward"
	s "api-merchant-backend/service/reward"

	"github.com/labstack/echo/v4"
)

func RewardCardCouponRoutes(echo *echo.Echo, conf config.Config) {
	db := config.ConnectDB(conf)

	repo := r.NewRewardRepository(db)
	service := s.NewRewardService(repo, conf)
	controllers := c.Controller{
		S: service,
	}

	e := echo.Group("/api")

	e.POST("/coupon", controllers.InsertRewardCard)
}
