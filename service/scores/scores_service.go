package scores

import "api-merchant-backend/entity"

type ScoresService interface {
	GetAllScores() ([]entity.ScoresResult, float64, error)
}
