package admin

import (
	"api-merchant-backend/entity"
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func NewMerchantRepository(DB *mongo.Database) MerchantRepository {
	return &Repository{
		DB: DB.Collection("merchant"),
	}
}

func (repo *Repository) InsertMerchant(merchant entity.MerchantCreate) error {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	merchant.Is_Registered = false
	merchant.Created_At = time.Now()

	defer cancel()
	_, err := repo.DB.InsertOne(context, merchant)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetAllMerchant() ([]entity.MerchantCreate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	cursor, err := repo.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.MerchantCreate
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *Repository) GetMerchantByID(id primitive.ObjectID) (*entity.MerchantCreate, error) {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	var result entity.MerchantCreate
	err := repo.DB.FindOne(context, filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, err
}

func (repo *Repository) UpdateMerchant(OldMerchant *entity.MerchantCreate, NewMerchant *entity.MerchantUpdate) (*entity.MerchantUpdate, error) {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	NewMerchant.Email = OldMerchant.Email
	NewMerchant.Updated_At = time.Now()
	filter := bson.M{
		"_id": OldMerchant.ID,
	}
	update := bson.M{
		"$set": NewMerchant,
	}
	_, err := repo.DB.UpdateOne(context, filter, update)
	if err != nil {
		return nil, err
	}
	return NewMerchant, nil
}

func (repo *Repository) DeleteMerchant(id primitive.ObjectID) error {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	_, err := repo.DB.DeleteOne(context, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) FindMerchantByEmail(email string) (*entity.MerchantCreate, error) {
	merchant := entity.MerchantCreate{}
	err := repo.DB.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&merchant)
	if err != nil {
		return nil, errors.New("wrong username")
	}
	return &merchant, nil
}
