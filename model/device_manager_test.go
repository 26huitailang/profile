package model

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"profile/config"
	"profile/database"
	"reflect"
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
			helperConfigInit()
			client, err := database.NewMongo()
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
			helperConfigInit()
			manager := initMongoManager(t)
			defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(item1)
			if err != nil {
				t.Fatal(err)
			}
			item2, err := manager.FindOne(bson.D{{"_id", insertResult.InsertedID}})
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, item2.BuyAt)
		})
	}
}

func TestDeviceManger_GetAllDevices(t *testing.T) {
	helperConfigInit()
	d1 := createOneDevice()
	d2 := createOneDevice()
	type fields struct {
		collection *mongo.Collection
		devices    []*Device
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// Add test cases.
		{name: "no item", fields: fields{collection: initCollection(t, "device"), devices: nil}, want: 0},
		{name: "one item", fields: fields{collection: initCollection(t, "device"), devices: []*Device{d1}}, want: 1},
		{name: "two item", fields: fields{collection: initCollection(t, "device"), devices: []*Device{d1, d2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DeviceManger{
				collection: tt.fields.collection,
			}
			defer helperDropCollection(m)

			if tt.fields.devices != nil {
				_, err := m.InsertMany(tt.fields.devices)
				if err != nil {
					log.Fatal(err)
				}
			}

			if got := len(m.GetAllDevices()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* helper func */
func helperDropCollection(m *DeviceManger) {
	m.DropCollection()
}

func helperConfigInit() {
	_ = os.Setenv("GO_ENV", "test")
	config.InitConfig()
}

func initMongoClient(t *testing.T) *mongo.Client {
	client, err := database.NewMongo()
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func initMongoManager(t *testing.T) *DeviceManger {
	client := initMongoClient(t)
	manager := NewDeviceManager(client, config.Cfg.Mongo.DB)
	return manager
}

func initCollection(t *testing.T, collectionName string) *mongo.Collection {
	client := initMongoClient(t)
	collection := client.Database(config.Cfg.Mongo.DB).Collection(collectionName)
	return collection
}

func createOneDevice() *Device {
	now := Now()
	item1 := NewDevice()
	item1.Name = "1"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = now
	return item1
}
