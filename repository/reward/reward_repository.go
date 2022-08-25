package reward

import (
	"api-merchant-backend/entity"
)

type RewardRepository interface {
	InsertRewardCard(coupon entity.RewardCardCoupon) (entity.RewardCardCoupon, error)
}
