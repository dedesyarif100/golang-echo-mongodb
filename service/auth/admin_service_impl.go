package auth

import (
	"api-merchant-backend/config"
	"api-merchant-backend/entity"
	"api-merchant-backend/repository/auth"
	"api-merchant-backend/utils"
	"errors"
)

type Service struct {
	repository auth.AdminRepository
	config     config.Config
}

func NewAdminService(repository auth.AdminRepository, config config.Config) AdminService {
	return &Service{
		repository: repository,
		config:     config,
	}
}

func (service *Service) Register(admin *entity.AuthRegister) (*entity.Admin, error) {
	result, _ := service.repository.FindAdminByEmail(admin.Email)
	if result != nil {
		return nil, errors.New("email is exists")
	}
	return service.repository.Register(admin)
}

func (service *Service) Login(admin *entity.AuthLogin) (*entity.ResponseLogin, error) {
	result, err := service.repository.FindAdminByEmail(admin.Email)
	if err != nil {
		return nil, err
	}
	errpw := utils.VerifyPassword(result.Password, admin.Password)
	if errpw != nil {
		return nil, errors.New("wrong password")
	}
	token, err := service.repository.CreateToken(result)
	response := &entity.ResponseLogin{
		Admin: *result,
		Token: *token,
	}
	return response, nil
}

func (service *Service) FindAdminByEmail(email string) (*entity.Admin, error) {
	return service.repository.FindAdminByEmail(email)
}
