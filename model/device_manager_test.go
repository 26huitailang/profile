package model

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile/database"
	"testing"
)

const MongoTestDB = "test"

func TestAssetManger_InsertOne(t *testing.T) {
	_id := primitive.NewObjectID()
	item1 := Device{ID: _id, Name: "0", Description: "a", Price: 99, Category: ElectronicEquipment}
	testCases := []struct {
		name string
		data interface{}
		want interface{}
	}{
		{name: "assert one asset", data: item1, want: _id},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := database.NewMongo(database.MongoUsername, database.MongoPassword, database.MongoHost, MongoTestDB)
			if err != nil {
				t.Fatal(err)
			}
			manager := NewDeviceManager(client)
			//defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(&item1)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, insertResult.InsertedID, tt.want)
		})
	}
}

func helperDropCollection(m *DeviceManger) {
	m.DropCollection()
}
