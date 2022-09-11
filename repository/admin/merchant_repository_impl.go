package admin

import (
	"api-merchant-backend/entity"
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "rating", Value: -1}}).SetLimit(1)
	cursor, _ := repo.DB.Find(context, filter, opts)

	var result []entity.MerchantResult
	cursor.All(context, &result)

	regex := regexp.MustCompile("[0-9]+")

	email := strings.Index(result[0].Email, "@")
	sliceEmailBeforeSymbol := result[0].Email[:email]
	sliceEmailAfterSymbol := result[0].Email[email:]
	findNumber := regex.FindAllString(sliceEmailBeforeSymbol, -1)
	findIndexNumber := strings.Index(sliceEmailBeforeSymbol, findNumber[0])
	admin := sliceEmailBeforeSymbol[:findIndexNumber]
	numEmail, _ := strconv.Atoi(findNumber[0])
	numEmail = numEmail + 1

	picEmail := strings.Index(result[0].Pic_Email, "@")
	slicePicEmailBeforeSymbol := result[0].Pic_Email[:picEmail]
	slicePicEmailAfterSymbol := result[0].Pic_Email[picEmail:]
	findNumber = regex.FindAllString(slicePicEmailBeforeSymbol, -1)
	findIndexNumber = strings.Index(slicePicEmailBeforeSymbol, findNumber[0])
	admin = slicePicEmailBeforeSymbol[:findIndexNumber]
	numPicEmail, _ := strconv.Atoi(findNumber[0])
	numPicEmail = numPicEmail + 1

	merchant.Rating = result[0].Rating + 1
	merchant.Email = admin+""+strconv.Itoa(numEmail)+""+sliceEmailAfterSymbol 
	merchant.Pic_Email = admin+""+strconv.Itoa(numEmail)+""+slicePicEmailAfterSymbol
	merchant.Is_Registered = false
	merchant.Created_At = time.Now()

	defer cancel()
	_, err := repo.DB.InsertOne(context, merchant)
	if err != nil {
		return err
	}
	return nil
}


func (repo *Repository) GetAllMerchant() ([]entity.MerchantResult, float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// CRUD OPERATIONS
	// READ OPERATIONS
	// SPECIFY A QUERY
	// QUERY CONDITION ( $eq, $ne, $gt, $gte, $lt, $lte, $in, $nin, $exists ) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "name", Value: "admin"}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		

		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "name", Value: bson.D{{Key: "$eq", Value: "admin"}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------

		
		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "name", Value: bson.D{{Key: "$ne", Value: "admin"}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "updated_at", Value: bson.D{{Key: "$exists", Value: true}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "updated_at", Value: bson.D{{Key: "$exists", Value: false}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$eq", Value: 10}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		

		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$ne", Value: 8}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$gt", Value: 9}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$gte", Value: 9}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$lt", Value: 8}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$lte", Value: 8}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$in", Value: []int{8, 10}}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "rating", Value: bson.D{{Key: "$nin", Value: []int{8, 9}}}}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{
		// 			Key: "$and", Value: bson.A{
		// 				bson.D{{
		// 					Key: "rating", 
		// 					Value: bson.D{{
		// 						Key: "$gt", 
		// 						Value: 8,
		// 					}},
		// 				}},
		// 				bson.D{{
		// 					Key: "rating", 
		// 					Value: bson.D{{
		// 						Key: "$lt", 
		// 						Value: 12,
		// 					}},
		// 				}},
		// 			}},
		// 		}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"},
		// 	{Key: "logo_url", Value: "id"},
		// 	{Key: "address", Value: bson.D{{
		// 		Key: "$ne",
		// 		Value: "jkt",
		// 	}}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"},
		// 	{Key: "logo_url", Value: "id"},
		// 	{Key: "$and", Value: bson.A{
		// 		bson.M{"address": "jkt"},
		// 		bson.D{{
		// 			Key: "rating",
		// 			Value: bson.D{{
		// 				Key: "$lt",
		// 				Value: 12,
		// 			}},
		// 		}},
		// 	},
		// }}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{"information.category": "food"}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{"information.category": bson.D{{
		// 	Key: "$in",
		// 	Value: []any{"food", "drink"},
		// }}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"}, 
		// 	{Key: "rating", Value: bson.D{{Key: "$lte", Value: 12}}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"}, 
		// 	{Key: "rating", Value: bson.D{{Key: "$gte", Value: 12}}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"}, 
		// 	{Key: "$and", Value: bson.A{
		// 		bson.D{{
		// 			Key: "rating", 
		// 			Value: bson.D{{
		// 				Key: "$gt", 
		// 				Value: 11,
		// 			}},
		// 		}},
		// 		bson.D{{
		// 			Key: "rating",
		// 			Value: bson.D{{
		// 				Key: "$lt",
		// 				Value: 14,
		// 			}},
		// 		}},
		// 	}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{"information.category": bson.D{{
		// 	Key: "$in",
		// 	Value: []any{"food", "drink"},
		// }}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{"information.category": bson.D{{
		// 	Key: "$nin",
		// 	Value: []any{"food", "drink"},
		// }}}
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"},
		// 	{Key: "rating", Value: bson.D{{Key: "$lt", Value: 12}}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.A{
		// 	bson.M{
		// 		"$lookup": bson.M{
		// 			"from":         "transaction",
		// 			"localField":   "email",
		// 			"foreignField": "name",
		// 			"as":           "trans",
		// 		},
		// 	},
		// 	bson.M{
		// 		"$unwind": "$trans",
		// 	},
		// 	bson.M{
		// 		"$lookup": bson.M{
		// 			"from":         "transaction_history",
		// 			"localField":   "trans.transaction_number",
		// 			"foreignField": "code",
		// 			"as":           "transaction_history",
		// 		},
		// 	},
		// 	bson.M{
		// 		"$unwind": "$transaction_history",
		// 	},
		// }
		// fmt.Println(ctx)
		// cursor, _ := repo.DB.Aggregate(ctx, filter)
		// var results []entity.MerchantResult
		// err := cursor.All(ctx, &results)
		// // count, err := repo.DB.CountDocuments(ctx, filter)
		// if err != nil {
		// 	return nil, 0, err
		// }
		// return results, 0, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "name", Value: "admin"},
		// 	{Key: "$and", Value: bson.A{
		// 		bson.D{{
		// 			Key: "rating", 
		// 			Value: bson.D{{
		// 				Key: "$gt", 
		// 				Value: 10,
		// 			}},
		// 		}},
		// 		bson.D{{
		// 			Key: "address",
		// 			Value: bson.D{{
		// 				Key: "$ne",
		// 				Value: "jkt",
		// 			}},
		// 		}},
		// 	}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.M{
		// 	"name": "merchant",
		// 	"address": "jkt",
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{
		// 	"name": "merchant",
		// 	"address": bson.D{{
		// 		Key: "$eq",
		// 		Value: "jkt",
		// 	}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.M{
		// 	"name": "merchant",
		// 	"address": bson.D{{
		// 		Key: "$ne",
		// 		Value: "jkt",
		// 	}},
		// }
		// cursor, _ := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// QUERY REGEX ( $regex ) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{{
		// 	Key: "name", 
		// 	Value: bson.D{{
		// 		Key: "$regex", 
		// 		Value:  "^m",
		// 	}},
		// }}
		// cursor, err := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{
		// 	Key: "logo_url", 
		// 	Value: bson.D{{
		// 		Key: "$regex", 
		// 		Value:  "^i",
		// 	}},
		// }}
		// cursor, err := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// QUERY ALL ( $all ) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// city := "sby"
		// filter := bson.D{
		// 	{
		// 		Key: "address", 
		// 		Value: bson.D{{
		// 			Key: "$all", 
		// 			Value: bson.A{city},
		// 		}},
		// 	},
		// }
		// cursor, err := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// COUNT DOCUMENTS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{{
		// 	Key: "rating", 
		// 	Value: bson.D{{
		// 		Key: "$gte", 
		// 		Value: 10,
		// 	}},
		// }}
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// var results []entity.MerchantResult
		// // fmt.Println("CEK VALUE :", count)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// matchStage := bson.D{{
		// 	Key: "$match", 
		// 	Value: bson.D{{
		// 		Key: "rating", 
		// 		Value: bson.D{{
		// 			Key: "$gt", 
		// 			Value: 10,
		// 		}},
		// 	}},
		// }}
		// countStage := bson.D{{
		// 	Key: "$count", 
		// 	Value: "rating",
		// }}
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{matchStage, countStage})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, results[0].Rating, err
		// }
		// return nil, results[0].Rating, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{
		// 	Key: "$count", 
		// 	Value: "rating",
		// }}
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{filter})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, 0, err
		// }
		// return nil, results[0].Rating, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// count, err := repo.DB.EstimatedDocumentCount(ctx)
		// // fmt.Println("CEK VALUE :", count)
		// val := float64(count)
		// var results []entity.MerchantResult
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// SET LIMIT ( $limit ) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetLimit(2)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// err = cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetLimit(2).SetSkip(3)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetSkip(2)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetSort(bson.D{{Key: "rating", Value: -1}}).SetLimit(1).SetSkip(1)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetSort(bson.D{{Key: "rating", Value: -1}}).SetLimit(1).SetSkip(2)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "$limit", Value: 5}}
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{filter})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.EstimatedDocumentCount(ctx)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// RETRIEVE DATA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// var results []entity.MerchantResult
		// filter := bson.D{{
		// 	Key: "$and", 
		// 	Value: bson.A{
		// 		bson.D{{
		// 			Key: "rating", 
		// 			Value: bson.D{{
		// 				Key: "$gte", 
		// 				Value: 8,
		// 			}},
		// 		}},
		// 		bson.D{{
		// 			Key: "rating", 
		// 			Value: bson.D{{
		// 				Key: "$lte", 
		// 				Value: 14,
		// 			}},
		// 		}},
		// 		bson.D{{
		// 			Key: "name",
		// 			Value: "admin",
		// 		}},
		// 		bson.D{{
		// 			Key: "address",
		// 			Value: "sby",
		// 		}},
		// 	},
		// }}
		// projection := bson.D{
		// 	{Key: "rating", Value: true},
		// 	{Key: "name", Value: true}, 
		// 	{Key: "email", Value: true}, 
		// 	{Key: "pic_email", Value: true}, 
		// }
		// opts := options.Find().SetProjection(projection)
		// cursor, _ := repo.DB.Find(ctx, filter, opts)
		// cursor.All(context.TODO(), &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// sort := bson.D{{Key: "rating", Value: -1}}
		// projection := bson.D{
		// 	{Key: "rating", Value: true}, 
		// 	{Key: "name", Value: true}, 
		// 	{Key: "email", Value: true}, 
		// 	{Key: "pic_email", Value: true}, 
		// }
		// opts := options.FindOne().SetSort(sort).SetProjection(projection)
		// var result bson.D
		// var results []entity.MerchantResult
		// var data entity.MerchantResult
		// repo.DB.FindOne(ctx, filter, opts).Decode(&result)
		// data.ID = result[0].Value.(primitive.ObjectID)
		// // fmt.Println(result)
		// data.Rating = float64(result[1].Value.(int32))
		// data.Name = result[2].Value.(string)
		// data.Email = result[3].Value.(string)
		// data.Pic_Email = result[4].Value.(string)
		// results = append(results, data)
		// return results, 0, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// groupStage := bson.D{
		// 	{Key: "$group", Value: bson.D{
		// 		{Key: "_id", Value: "$type"},
		// 		{Key: "rata-rata", Value: bson.D{
		// 			{Key: "$avg", Value: "$rating"},
		// 		}},
		// 	}},
		// }
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{groupStage})
		// var result []bson.M
		// cursor.All(ctx, &result)
		// // for _, element := range result {
		// // 	for key, val := range element {
		// // 		fmt.Println(key)
		// // 		fmt.Println(val)
		// // 	}
		// // 	// fmt.Println(element["_id"])
		// // 	// fmt.Println(element["rata-rata"])
		// // }
		// // fmt.Println(result[0]["rata-rata"])
		// var results []entity.MerchantResult
		// val := result[0]["rata-rata"].(float64)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// ACCESS DATA FROM A CURSOR >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		
	// ---------------------------------------------------------------------------------------------


	// RETRIEVE DISTINCH VALUES >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// result, err := repo.DB.Distinct(ctx, "name", bson.D{})
		// fmt.Println(result...)
	// ---------------------------------------------------------------------------------------------


	// SORT RESULTS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------

		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// SKIP RETURNED RESULTS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetSort(bson.D{{Key: "rating", Value: 1}}).SetSkip(2)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "rating", Value: -1}}}}
		skipStage := bson.D{{Key: "$skip", Value: 18}}
		cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{sortStage, skipStage})
		var results []entity.MerchantResult
		cursor.All(ctx, &results)
		if err != nil {
			return nil, float64(len(results)), err
		}
		return results, float64(len(results)), nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// LIMIT THE NUMBER OF RETURNED RESULTS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetLimit(2)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{}
		// opts := options.Find().SetSort(bson.D{{Key: "rating", Value: -1}}).SetLimit(2).SetSkip(1)
		// // opts := options.Find().SetSort(bson.D{{"rating", -1}}).SetSkip(1).SetLimit(2)
		// // opts := options.Find().SetLimit(2).SetSort(bson.D{{"rating", -1}}).SetSkip(1)
		// // opts := options.Find().SetLimit(2).SetSkip(1).SetSort(bson.D{{"rating", -1}})
		// // opts := options.Find().SetSkip(1).SetSort(bson.D{{"rating", -1}}).SetLimit(2)
		// // opts := options.Find().SetSkip(1).SetLimit(2).SetSort(bson.D{{Key: "rating", Value: -1}})

		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "$limit", Value: 3}}
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{filter})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ---------------------------------------------------------------------
	// ---------------------------------------------------------------------------------------------


	// SPECIFY WHICH FIELDS TO RETURN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------------------


	// SEARCH TEXT >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------
		// ---------------------------------------------------------------------
		// model := mongo.IndexModel{Keys: bson.D{{Key: "name", Value: "text"}}}
		// name, err := repo.DB.Indexes().CreateOne(ctx, model)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("NAME OF INDEX CREATED :", name)
		// var results []entity.MerchantResult
		// count, err := repo.DB.EstimatedDocumentCount(ctx)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------


		// ---------------------------------------------------------------------
		// model := mongo.IndexModel{Keys: bson.D{{Key: "address", Value: "text"}}}
		// name, err := repo.DB.Indexes().CreateOne(ctx, model)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println("NAME OF INDEX CREATED :", name)
		// var results []entity.MerchantResult
		// count, err := repo.DB.EstimatedDocumentCount(ctx)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		

		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "\"war\""}}}}
		// cursor, err := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		
		
		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "city"}}}}
		// cursor, err := repo.DB.Find(ctx, filter)
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		

		// ---------------------------------------------------------------------
		// var result bson.A
		// cursor.All(ctx, &result)
		// fmt.Println(result[0])

		// var result []bson.D
		// cursor.All(ctx, &result)
		// fmt.Println(result[0][1].Key)
		// fmt.Println(result[0][1].Value)
 
		// var result []bson.M
		// cursor.All(ctx, &result)
		// fmt.Println(result)
		// fmt.Println(result[1]["_id"])
		// ------------------------------------------------------------------------
		

		// ---------------------------------------------------------------------
		// filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "war"}}}}
		// sort := bson.D{
		// 	{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}},
		// }
		// projection := bson.D{
		// 	{Key: "_id", Value: true},
		// 	{Key: "rating", Value: true}, 
		// 	{Key: "name", Value: true}, 
		// 	{Key: "email", Value: true},
		// 	{Key: "pic_email", Value: true},
		// 	{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}, 
		// }
		// opts := options.Find().SetSort(sort).SetProjection(projection)
		// cursor, err := repo.DB.Find(ctx, filter, opts)
		// // var results []bson.D
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// count, err := repo.DB.CountDocuments(ctx, filter)
		// val := float64(count)
		// if err != nil {
		// 	return nil, val, err
		// }
		// return results, val, nil
		// ---------------------------------------------------------------------
		

		// ------------------------------------------------------------------------
		// filter := bson.D{
		// 	{Key: "$match", Value: bson.D{
		// 		{Key: "$text", Value: bson.D{
		// 			{Key: "$search", Value: "city"},
		// 		},
		// 	},
		// }}}
		// cursor, err := repo.DB.Aggregate(ctx, mongo.Pipeline{filter})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ------------------------------------------------------------------------


		// ------------------------------------------------------------------------
		// matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "city"}}}}}}
		// sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore" }}}}}}
		// projectStage := bson.D{
		// 	{Key: "$project", Value: bson.D{
		// 			{Key: "_id", Value: true},
		// 			{Key: "rating", Value: true}, 
		// 			{Key: "name", Value: true}, 
		// 			{Key: "email", Value: true}, 
		// 			{Key: "pic_email", Value: true}, 
		// 			{Key: "score", Value: bson.D{
		// 					{Key: "$meta", Value: "textScore"},
		// 				},
		// 			}, 
		// 		},
		// 	},
		// }
		// cursor, err := repo.DB.Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage, projectStage})
		// var results []entity.MerchantResult
		// cursor.All(ctx, &results)
		// if err != nil {
		// 	return nil, float64(len(results)), err
		// }
		// return results, float64(len(results)), nil
		// ------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------------------


	// WATCH FOR CHANGES >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// ---------------------------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------------------



	// findOptions := options.Find().SetProjection(bson.M{"_id": 1, "rating": 1, "name": 1})
	// findOptions.SetLimit(2)
	// cursor, err := repo.DB.Find(ctx, filter)
	// if err != nil {
	// 	return nil, err
	// }
	// return results, nil
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
	err := repo.DB.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&merchant)
	if err != nil {
		return nil, errors.New("wrong username")
	}
	return &merchant, nil
}
