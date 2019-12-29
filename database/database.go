package database

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func NewDB(filename string) (*gorm.DB, func()) {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic("连接数据库失败")
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	return db, func() {
		db.Close()
	}
}

const (
	MongoUsername string = "test"
	MongoPassword string = "123456"
	MongoHost     string = "localhost:27017"
	MongoDB       string = "develop"
)

func NewMongo(username, password, host, db string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := GenMongoURI(username, password, host, db)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

func GenMongoURI(username, password, host, db string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", username, password, host, db)
	return uri
}

func PingMongo(client *mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}
