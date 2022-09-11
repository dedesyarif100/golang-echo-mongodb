package server

import (
	"api-merchant-backend/config"
	"api-merchant-backend/routers"
	"api-merchant-backend/entity/validate"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	app.Validator = validate.NewValidationUtil()
	// koneksi database
	conf := config.Config{}

	// kumpulan routes
	routers.RewardCardCouponRoutes(app, conf)
	routers.MerchantRoutes(app, conf)
	routers.AdminRouters(app, conf)
	routers.HistoryRoutes(app, conf)

	return app
}
