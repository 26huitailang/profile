package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"profile/database"
)

type DeviceManger struct {
	client     *mongo.Client
	collection *mongo.Collection
	db         *mongo.Database
}

func NewDeviceManager(client *mongo.Client) *DeviceManger {
	return &DeviceManger{
		client:     client,
		db:         client.Database(database.MongoDB),
		collection: client.Database(database.MongoDB).Collection("device"),
	}
}

// insert
func (m *DeviceManger) InsertOne(item *Device) (*mongo.InsertOneResult, error) {
	insertResult, err := m.collection.InsertOne(context.TODO(), item)
	return insertResult, err
}
func (m *DeviceManger) InsertMany(items []interface{}) (*mongo.InsertManyResult, error) {
	insertResult, err := m.collection.InsertMany(context.TODO(), items)
	return insertResult, err
}

// find
func (m *DeviceManger) FindOne(filter bson.D) (*Device, error) {
	var result *Device
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	return result, err
}
func (m *DeviceManger) Find(filter bson.D, options *options.FindOptions) []*Device {
	var result []*Device
	cur, err := m.collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem *Device
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, elem)
	}

	return result
}

// update
func (m *DeviceManger) UpdateOne(filter bson.D, update bson.D) *mongo.UpdateResult {
	updateResult, err := m.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

// delete
func (m *DeviceManger) DeleteMany(filter bson.D) *mongo.DeleteResult {
	deleteResult, err := m.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult
}

// drop
func (m *DeviceManger) DropCollection() {
	err := m.collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

// implement interface

func (m *DeviceManger) InsertOneDevice(item *Device) (*Device, error) {
	ret, err := m.InsertOne(item)
	if err != nil {
		return nil, err
	}
	return m.FindOne(bson.D{{"_id", ret.InsertedID}})
}

func (m *DeviceManger) GetAllDevices() []*Device {
	return m.Find(bson.D{}, options.Find())
}

func (m *DeviceManger) UpdateOneDevice(item *Device) (*Device, error) {
	panic("implement me")
}

func (m *DeviceManger) GetOneDevice(id uint) (*Device, error) {
	panic("implement me")
}
