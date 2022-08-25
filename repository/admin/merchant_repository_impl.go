package admin

import (
	"api-merchant-backend/entity"
	"context"
	"fmt"
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

func (repo *Repository) InsertMerchant(merchant entity.Merchant) (entity.Merchant, error) {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	merchant.Is_Registered = false
	// fmt.Println("CEK VALUE :", merchant.Is_Registered)
	// return merchant, nil
	defer cancel()
	_, err := repo.DB.InsertOne(context, merchant)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repo *Repository) GetAllMerchant() ([]entity.Merchant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	cursor, err := repo.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.Merchant
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *Repository) GetMerchantByID(id primitive.ObjectID) (*entity.Merchant, error) {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	fmt.Println("CEK VALUE :", id);

	filter := bson.M{
		"id": id,
	}

	var result entity.Merchant
	err := repo.DB.FindOne(context, filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, err
}

func (repo *Repository) UpdateMerchant(id primitive.ObjectID, data *entity.Merchant) (*entity.Merchant, error) {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{
		"id": id,
	}
	update := bson.M{
		"$set": data,
	}
	_, err := repo.DB.UpdateOne(context, filter, update)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *Repository) DeleteMerchant(id primitive.ObjectID) error {
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{
		"id": id,
	}

	_, err := repo.DB.DeleteOne(context, filter)
	if err != nil {
		return err
	}
	return nil
}
