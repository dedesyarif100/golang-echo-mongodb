package scores

import (
	"api-merchant-backend/config"
	"api-merchant-backend/entity"
	"api-merchant-backend/repository/scores"
)

type Service struct {
	repository scores.ScoresRepository
	config config.Config
}

func NewScoresService(repository scores.ScoresRepository, config config.Config) ScoresService {
	return &Service{
		repository: repository,
		config: config,
	}
}

func (service *Service) GetAllScores() ([]entity.ScoresResult, float64, error) {
	return service.repository.GetAllScores()
}
