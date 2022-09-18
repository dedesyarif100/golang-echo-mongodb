package scores

import (
	"fmt"
	"time"
	"context"
	"api-merchant-backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Collection
}

func NewScoresRepository(DB *mongo.Database) ScoresRepository {
	return &Repository{
		DB: DB.Collection("scores"),
	}
}

func (repo *Repository) GetAllScores() ([]entity.ScoresResult, float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()


	// ------------------------- $addFields -------------------------
		// docs := []interface{}{
		// 	bson.D{
		// 		{Key: "student", Value: "Dede"}, 
		// 		{Key: "homework", Value: bson.A{10, 5, 10}}, 
		// 		{Key: "quiz", Value: bson.A{10, 8}}, 
		// 		{Key: "extracredit", Value: 4},
		// 		{Key: "embedded_score", Value: bson.D{
		// 			{Key: "uts", Value: 2},
		// 			{Key: "uas", Value: 3},
		// 		}},
		// 	},
		// 	bson.D{
		// 		{Key: "student", Value: "Hend"}, 
		// 		{Key: "homework", Value: bson.A{4, 6, 4}}, 
		// 		{Key: "quiz", Value: bson.A{8, 8}}, 
		// 		{Key: "extracredit", Value: 5},
		// 		{Key: "embedded_score", Value: bson.D{
		// 			{Key: "uts", Value: 2},
		// 			{Key: "uas", Value: 3},
		// 		}},
		// 	},
		// }
		// _, err := repo.DB.InsertMany(context.TODO(), docs)

		// filter := bson.D{{Key: "$limit", Value: 3}}


		// Using Two $addFields Stages -----------------------------
		// Adding Fields to an Embedded Document -----------------------------
		// filter := bson.A{
		// 	bson.M{"$addFields": bson.M{
		// 			"totalHomework": bson.M{"$sum": "$homework"},
		// 			"totalQuiz": bson.M{"$sum": "$quiz"},
		// 		},
		// 	},
		// 	bson.M{"$addFields": bson.M{
		// 			"totalScore": bson.M{"$add": bson.A{"$totalHomework", "$totalQuiz", "$extracredit"}},
		// 		},
		// 	},
		// 	bson.M{"$addFields": bson.M{
		// 			"embedded_score.scores": bson.M{"$add": bson.A{"$embedded_score.uts", "$embedded_score.uas"}},
		// 		},
		// 	},
		// }


		// Overwriting an existing field -----------------------------
		// filter := bson.A{
		// 	bson.M{"$addFields": bson.M{
		// 			"_id": "$student",
		// 			"student": "tess",
		// 		},
		// 	},
		// }


		// Add Element to an Array -----------------------------
		// filter := bson.A{
		// 	bson.M{"$match": bson.M{
		// 			"student": "dede",
		// 		},
		// 	},
		// 	bson.M{"$addFields": bson.M{
		// 			"homework": bson.M{
		// 				"$concatArrays": bson.A{
		// 					bson.D{{"$homework", 7}},
		// 				},
		// 			},
		// 		},
		// 	},
		// }

		// cursor, err := repo.DB.Aggregate(ctx, filter)
		// var results []entity.ScoresResult
		// var res []bson.M
		// cursor.All(ctx, &results)

		// for _, val := range res{
		// 	fmt.Println("CEK VALUE :", val)
		// }
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil

	// ------------------------- $bucket -------------------------

	// ------------------------- $count -------------------------
	// filter := bson.A{
	// 	bson.M{"$match": bson.M{
	// 		"extracredit": bson.M{
	// 			"$lt": 4,
	// 		},
	// 	}},
	// 	bson.M{
	// 		"$count": "passing_scores",
	// 	},
	// }
	// cursor, err := repo.DB.Aggregate(ctx, filter)
	// var results []entity.ScoresResult
	// var res []bson.M
	// cursor.All(ctx, &results)

	// for _, val := range res{
	// 	fmt.Println("CEK VALUE :", val)
	// }
	// if err != nil {
	// 	return nil, float64(len(results)), err
	// }
	// return results, float64(len(results)), nil

	// ------------------------- $replaceWith -------------------------
	// filter := bson.A{
	// 	bson.M{"$replaceWith": bson.M{
	// 		"$mergeObjects": bson.A{
	// 			bson.M{"_id": "$_id"},
	// 			bson.M{"student": "$student"},
	// 			"$embedded_score",
	// 		},
	// 	}},
	// }

	filter := bson.A{
		bson.M{"$replaceWith": bson.M{
			"$ifNull": bson.A{
				bson.A{"$student", 2},
			},
		}},
	}
	cursor, err := repo.DB.Aggregate(ctx, filter)
	var results []entity.ScoresResult
	var res []bson.M
	cursor.All(ctx, &results)

	for _, val := range res{
		fmt.Println("CEK VALUE :", val)
	}
	if err != nil {
		return nil, float64(len(results)), err
	}
	return results, float64(len(results)), nil
}

