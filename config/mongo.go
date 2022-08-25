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
	URI string
}

func ConnectDB(conf Config) *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URL")
	client, er := mongo.NewClient(options.Client().ApplyURI(uri))
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
	cli := client.Database("clean_arch")
	return cli
}
