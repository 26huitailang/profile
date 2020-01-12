package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func CreateIndexes(coll *mongo.Collection, indexField string) (string, error) {
	indexName, err := coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{indexField, bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	return indexName, err
}
