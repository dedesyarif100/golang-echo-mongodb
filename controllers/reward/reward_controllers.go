package reward

import (
	"api-merchant-backend/entity"
	"api-merchant-backend/service/reward"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	S reward.RewardService
}

func (h *Controller) InsertRewardCard(c echo.Context) error {
	coupon := entity.RewardCardCoupon{
		ID: primitive.NewObjectID(),
		Coupon_Detail: entity.Coupon_Detail{
			ID: primitive.NewObjectID(),
		},
	}

	err := c.Bind(&coupon)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"messages": err.Error(),
		})
	}

	reward, er := h.S.InsertRewardCard(coupon)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"messages": er.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"messages": "successfully created reward card coupon",
		"data":     reward,
	})
}
