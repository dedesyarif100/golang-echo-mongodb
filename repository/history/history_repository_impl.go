package history

import (
	"time"
	"context"
	"api-merchant-backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func NewHistoryRepository(DB *mongo.Database) HistoryRepository {
	return &Repository{
		DB: DB.Collection("history_point_test"),
	}
}

func (repo *Repository) GetCountPointByIdCoupon() (*entity.ResponseHistory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// filter := bson.M{
	// 	"$group": bson.M{
	// 		"_id": "",
	// 		"count": bson.M{"$sum": "$quantity"}, 
	// 	},
	// }

	filter := bson.M{
		"$match": bson.M{
			"status": "success",
		},
		"$sum": "point_amount",
	}

	cursor, err := repo.DB.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results entity.ResponseHistory
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}