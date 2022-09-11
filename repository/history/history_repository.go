package history

import (
	"api-merchant-backend/entity"
)

type HistoryRepository interface {
	GetCountPointByIdCoupon() (*entity.ResponseHistory, error)
}
