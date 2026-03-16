package util

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error env file")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	DB = client.Database(dbName)

	log.Println("Connected")
}
