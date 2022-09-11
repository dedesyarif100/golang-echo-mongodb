package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryPoint struct {
	ID          primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Coupon_ID   primitive.ObjectID	`json:"coupon_id" bson:"coupon_id"`
	Coupon      Coupon_Detail		`json:"coupon" bson:"coupon"`
	PointCoupon int					`json:"point_coupon" bson:"point_coupon"`
	PointAmount int					`json:"point_amount" bson:"point_amount"`
	UserID      int					`json:"user_id" bson:"user_id"`
	Phone       string				`json:"phone" bson:"phone"`
	StatusPoint string				`json:"status_point" bson:"status_point"`
	Created_At  time.Time			`json:"created_at" bson:"created_at"`
	Amount      int					`json:"amount" bson:"amount"`
	Status      string				`json:"status" bson:"status"`
	ReceiptNo   string				`json:"receipt_no" bson:"receipt_no"`
}

type ResponseHistory struct {
	Count		int		`json:"count" bson:"count"`
}