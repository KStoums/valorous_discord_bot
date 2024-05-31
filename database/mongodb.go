package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
)

var MongoDb *mongo.Database

func StartMongoDb() error {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	// TODO FIX THAT TO DISCONNECT MONGO CLIENT
	//defer mongoClient.Disconnect(ctx)

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoDb = mongoClient.Database(strings.ToLower(os.Getenv("PROJECT_NAME")))
	return nil
}
