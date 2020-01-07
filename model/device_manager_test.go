package model

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"profile/config"
	"profile/database"
	"testing"
)

func TestAssetManger_InsertOne(t *testing.T) {
	item1 := NewDevice()
	item1.Name = "0"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = Now()

	testCases := []struct {
		name string
		data *Device
		want interface{}
	}{
		{name: "assert one asset", data: item1, want: item1.ID},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := database.NewMongo(config.Cfg.Mongo.Username, config.Cfg.Password, config.Cfg.Host, config.Cfg.Mongo.DB)
			if err != nil {
				t.Fatal(err)
			}
			manager := NewDeviceManager(client, config.Cfg.Mongo.DB)
			defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(item1)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, insertResult.InsertedID)
		})
	}
}

func TestAssetManger_InsertOne_Time_OK(t *testing.T) {
	now := Now()
	item1 := NewDevice()
	item1.Name = "1"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = now

	testCases := []struct {
		name string
		data *Device
		want Timestamp
	}{
		{name: "insert date with millisecond accuracy", data: item1, want: now},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := database.NewMongo(config.Cfg.Mongo.Username, config.Cfg.Password, config.Cfg.Host, config.Cfg.Mongo.DB)
			if err != nil {
				t.Fatal(err)
			}
			manager := NewDeviceManager(client, config.Cfg.Mongo.DB)
			defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(item1)
			if err != nil {
				t.Fatal(err)
			}
			item2, err := manager.FindOne(bson.D{{"_id", insertResult.InsertedID}})
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("want: %v got: %v", tt.want.String(), item2.BuyAt.String())
			assert.Equal(t, tt.want, item2.BuyAt)
		})
	}
}
func helperDropCollection(m *DeviceManger) {
	m.DropCollection()
}
