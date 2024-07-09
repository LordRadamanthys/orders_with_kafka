package mongodbconfig

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Collection

func InitMongoDBConfig() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:123@localhost:27017/"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	MongoClient = client.Database("OrderDB").Collection("Orders")
}
