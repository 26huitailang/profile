package model

import (
	"github.com/stretchr/testify/assert"
	"profile/database"
	"testing"
	"time"
)

const MongoTestDB = "test"
const MongoTestUsername = "test"
const MongoTestPassword = "test"

func TestAssetManger_InsertOne(t *testing.T) {
	item1 := NewDevice()
	item1.Name = "0"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = Timestamp(time.Now())

	testCases := []struct {
		name string
		data *Device
		want interface{}
	}{
		{name: "assert one asset", data: item1, want: item1.ID},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := database.NewMongo(MongoTestUsername, MongoTestPassword, database.MongoHost, MongoTestDB)
			if err != nil {
				t.Fatal(err)
			}
			manager := NewDeviceManager(client, MongoTestDB)
			//defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(item1)
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
