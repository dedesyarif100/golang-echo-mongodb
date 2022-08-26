package auth

import (
	"api-merchant-backend/entity"
)

type AdminService interface {
	Register(admin *entity.AuthRegister) (*entity.Admin, error)
	Login(admin *entity.AuthLogin) (*entity.ResponseLogin, error)
	FindAdminByEmail(email string) (*entity.Admin, error)
}
