package history

import (
	"api-merchant-backend/config"
	"api-merchant-backend/entity"
	"api-merchant-backend/repository/history"
)

type Service struct {
	repository history.HistoryRepository
	config config.Config
}

func NewHistoryService(repository history.HistoryRepository, config config.Config) HistoryService {
	return &Service{
		repository: repository,
		config: config,
	}
}

func (service *Service) GetCountPointByIdCoupon() (*entity.ResponseHistory, error) {
	return service.repository.GetCountPointByIdCoupon()
}