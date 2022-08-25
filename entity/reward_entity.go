package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RewardCardCoupon struct {
	ID             primitive.ObjectID `json:"id" bson:"id"`
	URL_Logo       string             `json:"url_logo" bson:"url_logo"`
	Title          string             `json:"title" bson:"title"`
	Card_No        string             `json:"card_no" bson:"card_no"`
	Start_At       string             `json:"start_at" bson:"start_at"`
	End_At         string             `json:"end_at" bson:"end_at"`
	Status         string             `json:"status" bson:"status"`
	Card_Claimed   int                `json:"card_claimed" bson:"card_claimed"`
	Points_Claimed int                `json:"points_claimed" bson:"points_claimed"`
	TnC            string             `json:"tnc" bson:"tnc"`
	Points_Detail  Points_Detail      `json:"points_detail" bson:"points_detail"`
	Coupon_Detail  Coupon_Detail      `json:"coupon_detail" bson:"coupon_detail"`

	Created_At *time.Time `json:"created_at" bson:"created_at"`
}

type Points_Detail struct {
	Amount        int    `json:"amount" bson:"amount"`
	Options_Issue string `json:"options_issue" bson:"options_issue"`
	Total_Points  int    `json:"total_points" bson:"total_points"`
}

type Coupon_Detail struct {
	ID              primitive.ObjectID `json:"id" bson:"id"`
	Title           string             `json:"title" bson:"title"`
	Item            string             `json:"item" bson:"item"`
	Status          string             `json:"status" bson:"status"`
	Total_Points    int                `json:"total_points" bson:"total_points"`
	Expired_At      string             `json:"expired_at" bson:"expired_at"`
	Coupon_Claimed  int                `json:"coupon_claimed" bson:"coupon_claimed"`
	Coupon_Redeemed int                `json:"coupon_redeemed" bson:"coupon_redeemed"`
	TnC             string             `json:"tnc" bson:"tnc"`

	Created_At *time.Time `json:"created_at" bson:"created_at"`
}
