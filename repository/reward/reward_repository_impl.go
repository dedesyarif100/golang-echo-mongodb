package reward

import (
	"api-merchant-backend/entity"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	DB *mongo.Collection
}

func NewRewardRepository(DB *mongo.Database) RewardRepository {
	return &repository{
		DB: DB.Collection("rewards_coupon"),
	}
}

func (r *repository) InsertRewardCard(coupon entity.RewardCardCoupon) (entity.RewardCardCoupon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.DB.InsertOne(ctx, coupon)
	if err != nil {
		return coupon, err
	}

	return coupon, nil
}
