package server

import (
	"api-merchant-backend/config"
	"api-merchant-backend/routers"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	// koneksi database
	conf := config.Config{}

	// kumpulan routes
	routers.RewardCardCouponRoutes(app, conf)
	routers.MerchantRoutes(app, conf)

	return app
}
