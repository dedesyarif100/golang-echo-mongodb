package reward

import (
	"api-merchant-backend/config"
	"api-merchant-backend/entity"
	"api-merchant-backend/repository/reward"
)

type Service struct {
	repo reward.RewardRepository
	conf config.Config
}

func NewRewardService(repo reward.RewardRepository, conf config.Config) RewardService {
	return &Service{
		repo: repo,
		conf: conf,
	}
}

func (s *Service) InsertRewardCard(coupon entity.RewardCardCoupon) (entity.RewardCardCoupon, error) {
	return s.repo.InsertRewardCard(coupon)
}
