package history

import "api-merchant-backend/entity"

type HistoryService interface {
	GetCountPointByIdCoupon() (*entity.ResponseHistory, error)
}