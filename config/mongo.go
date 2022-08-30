package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MONGO_URL    string
	MONGO_DBNAME string
	JWT_SECRET	 string
}

func ConnectDB(conf Config) *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dburl := os.Getenv("MONGO_URL")
	dbname := os.Getenv("MONGO_DBNAME")
	dbport := os.Getenv("MONGO_PORT")
	url := "mongodb://" + dburl + ":" + dbport

	client, er := mongo.NewClient(options.Client().ApplyURI(url))
	if er != nil {
		log.Fatal(er)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	er = client.Connect(ctx)
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println("Connected to MongoDB")
	cli := client.Database(dbname)
	return cli
}
