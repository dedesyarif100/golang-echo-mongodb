package admin

import (
	"api-merchant-backend/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MerchantRepository interface {
	InsertMerchant(merchant entity.MerchantCreate) (entity.MerchantCreate, error)
	GetAllMerchant() ([]entity.MerchantCreate, error)
	GetMerchantByID(ID primitive.ObjectID) (*entity.MerchantCreate, error)
	UpdateMerchant(ID primitive.ObjectID, merchant *entity.MerchantUpdate) (*entity.MerchantUpdate, error)
	DeleteMerchant(ID primitive.ObjectID) error
	FindMerchantByEmail(email string) (*entity.MerchantCreate, error)
}