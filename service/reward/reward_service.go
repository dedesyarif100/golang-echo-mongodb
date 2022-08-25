package reward

import "api-merchant-backend/entity"

type RewardService interface {
	InsertRewardCard(coupon entity.RewardCardCoupon) (entity.RewardCardCoupon, error)
}
