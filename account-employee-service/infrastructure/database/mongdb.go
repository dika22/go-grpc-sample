package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB return mongo db read instance
func InitMongoDB(ctx context.Context) *mongo.Database {

	// init mongodb
	host := os.Getenv("MONGO_DB_HOST")
	database := os.Getenv("MONGO_DB_NAME")

	fmt.Println("HOST", host)
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		panic(err)
	}
	if err := client.Connect(ctx); err != nil {
		panic(err)
	}

	return client.Database(database)
}

