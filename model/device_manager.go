package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DeviceManger struct {
	collection *mongo.Collection
}

func NewDeviceManager(client *mongo.Client, db string) *DeviceManger {
	return &DeviceManger{
		collection: client.Database(db).Collection("device"),
	}
}

// insert
func (m *DeviceManger) InsertOne(item *Device) (*mongo.InsertOneResult, error) {
	insertResult, err := m.collection.InsertOne(context.TODO(), item)
	return insertResult, err
}
func (m *DeviceManger) InsertMany(items []*Device) (*mongo.InsertManyResult, error) {
	var data []interface{}
	for _, item := range items {
		data = append(data, item)
	}
	insertResult, err := m.collection.InsertMany(context.TODO(), data)
	return insertResult, err
}

// find
func (m *DeviceManger) FindOne(filter bson.D) (*Device, error) {
	result := &Device{}
	err := m.collection.FindOne(context.TODO(), filter).Decode(result)
	return result, err
}
func (m *DeviceManger) Find(filter bson.D, options *options.FindOptions) []*Device {
	var result []*Device
	cur, err := m.collection.Find(context.TODO(), filter, options)
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var elem Device
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		result = append(result, &elem)
	}
	if result == nil {
		result = []*Device{}
	}
	return result
}

// update
func (m *DeviceManger) UpdateOne(filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	updateResult, err := m.collection.UpdateOne(context.TODO(), filter, update)
	return updateResult, err
}

// delete
func (m *DeviceManger) DeleteMany(filter bson.D) (*mongo.DeleteResult, error) {
	deleteResult, err := m.collection.DeleteMany(context.TODO(), filter)
	return deleteResult, err
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
	filter := bson.D{{"_id", bson.D{{"$eq", item.ID}}}}
	item.UpdatedAt = Now()
	data, err := bson.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("bson.Marshal: %v", err)
	}
	var doc bson.D
	err = bson.Unmarshal(data, &doc)

	if err != nil {
		return nil, fmt.Errorf("bson.Unmarshal: %v", err)
	}

	update := bson.D{{"$set", doc}}

	_, err = m.UpdateOne(filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOne: %v", err)
	}

	return m.FindOne(bson.D{{"_id", item.ID}})
}

func (m *DeviceManger) GetOneDevice(id primitive.ObjectID) (*Device, error) {
	return m.FindOne(bson.D{{"_id", id}})
}

func (m *DeviceManger) DeleteDeviceList(ids []primitive.ObjectID) (ret *mongo.DeleteResult, err error) {
	var objectIDs bson.A
	for _, id := range ids {
		objectIDs = append(objectIDs, id)
	}
	ret, err = m.DeleteMany(bson.D{{"_id", bson.D{{"$in", objectIDs}}}})
	return
}
