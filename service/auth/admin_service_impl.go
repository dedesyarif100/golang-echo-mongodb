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

func (service *Service) Register(admin *entity.AuthRegister) error {
	result, _ := service.repository.FindAdminByEmail(admin.Email)
	if result != nil {
		return errors.New("email is exists")
	}
	hashpw, _ := utils.Hash(admin.Password)
	admin.Password = string(hashpw)
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
	
	token, _ := utils.GenerateAccessToken(result.ID, service.config.JWT_SECRET, result.Email)
	response := &entity.ResponseLogin{
		Admin: *result,
		Token: token,
	}
	return response, nil
}

func (service *Service) FindAdminByEmail(email string) (*entity.Admin, error) {
	return service.repository.FindAdminByEmail(email)
}
