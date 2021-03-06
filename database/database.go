package database

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"profile/config"
	"time"
)

func NewMongo() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := GenMongoURI(config.Cfg.Mongo.Username, config.Cfg.Mongo.Password, config.Cfg.Mongo.Host, config.Cfg.Mongo.DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

func GenMongoURI(username, password, host, db string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=%s", username, password, host, db, db)
	log.Printf("mongodb uri: %s", uri)
	return uri
}

func PingMongo(client *mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}
