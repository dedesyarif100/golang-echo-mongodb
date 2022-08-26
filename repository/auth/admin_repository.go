package auth

import (
	"api-merchant-backend/entity"
)

type AdminRepository interface {
	Register(admin *entity.AuthRegister) (*entity.Admin, error)
	FindAdminByEmail(email string) (*entity.Admin, error)
	CreateToken(admin *entity.Admin) (*string, error)
}
