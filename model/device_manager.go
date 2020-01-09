package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"profile/config"
)

type DeviceManger struct {
	BaseManager
}

func NewDeviceManager(client *mongo.Client) *DeviceManger {
	return &DeviceManger{
		BaseManager: BaseManager{
			collection: client.Database(config.Cfg.Mongo.DB).Collection("device"),
		},
	}
}

// implement interface
func (m *DeviceManger) InsertOneDevice(item *Device) (*Device, error) {
	ret, err := m.InsertOne(item)
	if err != nil {
		return nil, err
	}
	singleResult := m.FindOne(bson.D{{"_id", ret.InsertedID}})
	var device Device
	err = singleResult.Decode(&device)
	return &device, err
}

func (m *DeviceManger) GetAllDevices() ([]*Device, error) {
	result := make([]*Device, 0)

	cur, err := m.FindMany(bson.D{}, options.Find())
	if err != nil {
		err = fmt.Errorf("GetAllDevices: %v", err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem Device
		err := cur.Decode(&elem)
		if err != nil {
			err = fmt.Errorf("GetAllDevices: %v", err)
			return nil, err
		}
		result = append(result, &elem)
	}
	return result, nil
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

	return m.GetOneDevice(item.ID)
}

func (m *DeviceManger) GetOneDevice(id primitive.ObjectID) (*Device, error) {
	singleResult := m.FindOne(bson.D{{"_id", id}})
	var device Device
	err := singleResult.Decode(&device)
	return &device, err
}

func (m *DeviceManger) DeleteDeviceList(ids []primitive.ObjectID) (ret *mongo.DeleteResult, err error) {
	var objectIDs bson.A
	for _, id := range ids {
		objectIDs = append(objectIDs, id)
	}
	ret, err = m.DeleteMany(bson.D{{"_id", bson.D{{"$in", objectIDs}}}})
	return
}
