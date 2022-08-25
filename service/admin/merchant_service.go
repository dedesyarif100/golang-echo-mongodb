package admin

import (
	"api-merchant-backend/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MerchantService interface {
	GetAllMerchant() ([]entity.Merchant, error)
	InsertMerchant(merchant entity.Merchant) (entity.Merchant, error)
	GetMerchantByID(ID primitive.ObjectID) (*entity.Merchant, error)
	UpdateMerchant(ID primitive.ObjectID, merchant *entity.Merchant) (*entity.Merchant, error)
	DeleteMerchant(ID primitive.ObjectID) error
}