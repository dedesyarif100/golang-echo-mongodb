package admin

import (
	"api-merchant-backend/config"
	"api-merchant-backend/entity"
	"api-merchant-backend/repository/admin"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repository admin.MerchantRepository
	config config.Config
}

func NewMerchantService(repository admin.MerchantRepository, config config.Config) MerchantService {
	return &Service{
		repository: repository,
		config: config,
	}
}

func (service *Service) InsertMerchant(merchant entity.MerchantCreate) (entity.MerchantCreate, error) {
	result, _ := service.repository.FindMerchantByEmail(merchant.Email)
	if result != nil {
		return *result, errors.New("email is exists")
	}
	return service.repository.InsertMerchant(merchant)
}


func (service *Service) GetAllMerchant() ([]entity.MerchantCreate, error) {
	return service.repository.GetAllMerchant()
}


func (service *Service) GetMerchantByID(id primitive.ObjectID) (*entity.MerchantCreate, error) {
	return service.repository.GetMerchantByID(id)
}


func (service *Service) UpdateMerchant(id primitive.ObjectID, merchant *entity.MerchantUpdate) (*entity.MerchantUpdate, error) {
	_, err := service.repository.GetMerchantByID(id)
	if err != nil {
		return merchant, err
	}
	return service.repository.UpdateMerchant(id, merchant)
}

func (service *Service) DeleteMerchant(id primitive.ObjectID) error {
	_, err := service.repository.GetMerchantByID(id)
	if err != nil {
		return err
	}
	return service.repository.DeleteMerchant(id)
}

func (service *Service) FindMerchantByEmail(email string) (*entity.MerchantCreate, error) {
	return service.repository.FindMerchantByEmail(email)
}