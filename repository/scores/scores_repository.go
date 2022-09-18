package scores

import "api-merchant-backend/entity"

type ScoresRepository interface {
	GetAllScores() ([]entity.ScoresResult, float64, error)
}