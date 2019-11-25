package database

import (
	"context"
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

func NewMongo() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	return client, err
}

func PingMongo(client *mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}
