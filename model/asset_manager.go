package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type AssetManger struct {
	client     *mongo.Client
	collection *mongo.Collection
	db         *mongo.Database
}

func NewAssetManager(client *mongo.Client) *AssetManger {
	return &AssetManger{
		client:     client,
		db:         client.Database("mock"),
		collection: client.Database("mock").Collection("asset"),
	}
}

// insert
func (m *AssetManger) InsertOne(item *Asset) (*mongo.InsertOneResult, error) {
	insertResult, err := m.collection.InsertOne(context.TODO(), item)
	return insertResult, err
}
func (m *AssetManger) InsertMany(items []interface{}) (*mongo.InsertManyResult, error) {
	insertResult, err := m.collection.InsertMany(context.TODO(), items)
	return insertResult, err
}

// find
func (m *AssetManger) FindOne(filter bson.D) (*Asset, error) {
	var result *Asset
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	return result, err
}
func (m *AssetManger) Find(filter bson.D, options *options.FindOptions) []*Asset {
	var result []*Asset
	cur, err := m.collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem *Asset
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, elem)
	}

	return result
}

// update
func (m *AssetManger) UpdateOne(filter bson.D, update bson.D) *mongo.UpdateResult {
	updateResult, err := m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

// delete
func (m *AssetManger) DeleteMany(filter bson.D) *mongo.DeleteResult {
	deleteResult, err := m.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
}

// drop
func (m *AssetManger) DropCollection() {
	err := m.collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
