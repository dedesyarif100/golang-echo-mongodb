package auth

import (
	"api-merchant-backend/entity"
	"api-merchant-backend/utils"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	DB *mongo.Collection
}

func NewAdminRepository(DB *mongo.Database) AdminRepository {
	return &Repository{
		DB: DB.Collection("admin"),
	}
}

func (repo *Repository) Register(admin *entity.AuthRegister) (*entity.Admin, error) {
	findOptions := options.Find()
	cur, err := repo.DB.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var results []entity.AuthRegister
	for cur.Next(context.TODO()) {
		var elem entity.AuthRegister
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	hashpw, _ := utils.Hash(admin.Password)
	insertData := &entity.Admin{
		Username:   admin.Username,
		Email:      admin.Email,
		Password:   string(hashpw),
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
	_, err = repo.DB.InsertOne(context.Background(), &insertData)
	if err != nil {
		return nil, err
	}
	return insertData, nil
}

func (repo *Repository) CreateToken(Data *entity.Admin) (*string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &entity.Claims{
		Username: Data.Username,
		Email:    Data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := os.Getenv("SECRET_JWT")
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	return &token_jwt, err
}

func (repo *Repository) FindAdminByEmail(email string) (*entity.Admin, error) {
	admin := entity.Admin{}
	err := repo.DB.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&admin)
	if err != nil {
		return nil, errors.New("wrong Email")
	}
	return &admin, nil
}
