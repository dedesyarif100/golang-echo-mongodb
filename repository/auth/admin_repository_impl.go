package auth

import (
	"api-merchant-backend/entity"
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func NewAdminRepository(DB *mongo.Database) AdminRepository {
	return &Repository{
		DB: DB.Collection("admin"),
	}
}

func (repo *Repository) Register(admin *entity.AuthRegister) error {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	admin.Created_At = time.Now()
	defer cancel()

	_, err := repo.DB.InsertOne(context, admin)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) FindAdminByEmail(email string) (*entity.Admin, error) {
	admin := entity.Admin{}
	err := repo.DB.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&admin)
	if err != nil {
		return nil, errors.New("Admin Not Found")
	}
	return &admin, nil
}
