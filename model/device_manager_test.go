package model_test

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"profile/config"
	"profile/database"
	"profile/model"
	"reflect"
	"testing"
)

func TestAssetManger_InsertOne(t *testing.T) {
	item1 := model.NewDevice()
	item1.Name = "0"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = model.Now()

	testCases := []struct {
		name string
		data *model.Device
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
			manager := model.NewDeviceManager(client)
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
	now := model.Now()
	item1 := model.NewDevice()
	item1.Name = "1"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = now

	testCases := []struct {
		name string
		data *model.Device
		want model.Timestamp
	}{
		{name: "insert date with millisecond accuracy", data: item1, want: now},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			helperConfigInit()
			manager := initDeviceManager(t)
			defer helperDropCollection(manager)

			insertResult, err := manager.InsertOne(item1)
			if err != nil {
				t.Fatal(err)
			}

			if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
				item2, err := manager.GetOneDevice(oid)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.want, item2.BuyAt)
			}
		})
	}
}

func TestDeviceManger_GetAllDevices(t *testing.T) {
	helperConfigInit()
	d1 := createOneDevice()
	d2 := createOneDevice()
	type fields struct {
		devices []*model.Device
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// Add test cases.
		{name: "no item", fields: fields{devices: nil}, want: 0},
		{name: "one item", fields: fields{devices: []*model.Device{d1}}, want: 1},
		{name: "two item", fields: fields{devices: []*model.Device{d1, d2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := initDeviceManager(t)
			defer helperDropCollection(m)

			if tt.fields.devices != nil {
				var data []interface{}
				for _, item := range tt.fields.devices {
					data = append(data, item)
				}
				_, err := m.InsertMany(data)
				if err != nil {
					log.Fatal(err)
				}
			}

			ret, _ := m.GetAllDevices()
			if got := len(ret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceManger_UpdateOneDevice(t *testing.T) {
	helperConfigInit()
	device1 := model.NewDevice()
	type fields struct {
	}
	type args struct {
		item *model.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Device
		wantErr bool
	}{
		// Add test cases.
		{name: "update ok", args: args{item: device1}, want: device1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := initDeviceManager(t)
			defer helperDropCollection(m)
			device1.Price = 66
			_, err := m.InsertOne(device1)
			if err != nil {
				t.Errorf("Insert error: %v", err)
			}
			device1.Price = 99
			got, err := m.UpdateOneDevice(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateOneDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Price, got.Price)
		})
	}
}

func TestDeviceManger_DeleteDeviceList(t *testing.T) {
	helperConfigInit()
	device1 := model.NewDevice()
	device2 := model.NewDevice()

	type fields struct {
		collection *mongo.Collection
		devices    []*model.Device
	}
	type args struct {
		ids []primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet *mongo.DeleteResult
		wantErr bool
	}{
		// Add test cases.
		{name: "delete 2 ok", fields: fields{}, args: args{ids: []primitive.ObjectID{device1.ID, device2.ID}}, wantRet: &mongo.DeleteResult{DeletedCount: 2}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := initDeviceManager(t)
			defer helperDropCollection(m)

			_, err := m.InsertMany([]interface{}{device1, device2})
			devices, _ := m.GetAllDevices()
			assert.Equal(t, tt.wantRet.DeletedCount, int64(len(devices)))

			gotRet, err := m.DeleteDeviceList(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDeviceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("DeleteDeviceList() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

/* helper func */
type dropper interface {
	DropCollection()
}

func helperDropCollection(m dropper) {
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

func initDeviceManager(t *testing.T) *model.DeviceManger {
	client := initMongoClient(t)
	manager := model.NewDeviceManager(client)
	return manager
}

func initCollection(t *testing.T, collectionName string) *mongo.Collection {
	client := initMongoClient(t)
	collection := client.Database(config.Cfg.Mongo.DB).Collection(collectionName)
	return collection
}

func createOneDevice() *model.Device {
	now := model.Now()
	item1 := model.NewDevice()
	item1.Name = "1"
	item1.Description = "desc"
	item1.Price = 99
	item1.Category = 1
	item1.BuyAt = now
	return item1
}
