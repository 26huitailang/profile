package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	MongoUsername string = "develop"
	MongoPassword string = "develop"
	MongoHost     string = "10.200.233.224:27017"
	MongoDB       string = "develop"
)

func NewMongo(username, password, host, db string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := GenMongoURI(username, password, host, db)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

func GenMongoURI(username, password, host, db string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=develop", username, password, host, db)
	return uri
}

func PingMongo(client *mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}
