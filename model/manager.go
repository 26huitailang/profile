package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewManager(deviceManger *DeviceManger) *Manager {
	return &Manager{DeviceManger: deviceManger}
}

type IModelManager interface {
	InsertOne(item interface{}) (*mongo.InsertOneResult, error)
	InsertMany(items []interface{}) (*mongo.InsertManyResult, error)
	FindOne(filter bson.D) *mongo.SingleResult
	FindMany(filter bson.D, options *options.FindOptions) (*mongo.Cursor, error)
	UpdateOne(filter bson.D, update bson.D) (*mongo.UpdateResult, error)
	DeleteMany(filter bson.D) (*mongo.DeleteResult, error)
	DropCollection()
}

type Manager struct {
	*DeviceManger
}

type BaseManager struct {
	collection *mongo.Collection
}

// insert
func (m *BaseManager) InsertOne(item interface{}) (*mongo.InsertOneResult, error) {
	insertResult, err := m.collection.InsertOne(context.TODO(), item)
	return insertResult, err
}

func (m *BaseManager) InsertMany(items []interface{}) (*mongo.InsertManyResult, error) {
	insertResult, err := m.collection.InsertMany(context.TODO(), items)
	return insertResult, err
}

// find
func (m *BaseManager) FindOne(filter bson.D) *mongo.SingleResult {
	singleResult := m.collection.FindOne(context.TODO(), filter)
	return singleResult
}

func (m *BaseManager) FindMany(filter bson.D, options *options.FindOptions) (*mongo.Cursor, error) {
	cur, err := m.collection.Find(context.TODO(), filter, options)
	if err != nil {
		err = fmt.Errorf("FindMany: %v", err)
	}
	return cur, err
}

// update
func (m *BaseManager) UpdateOne(filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	updateResult, err := m.collection.UpdateOne(context.TODO(), filter, update)
	return updateResult, err
}

// delete
func (m *BaseManager) DeleteMany(filter bson.D) (*mongo.DeleteResult, error) {
	deleteResult, err := m.collection.DeleteMany(context.TODO(), filter)
	return deleteResult, err
}

// drop
func (m *BaseManager) DropCollection() {
	err := m.collection.Drop(context.TODO())
	if err != nil {
		panic(err)
	}
}
