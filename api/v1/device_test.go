package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"profile/api"
	v1 "profile/api/v1"
	"profile/model"
	"reflect"
	"testing"
	"time"
)

// todo: unittest for Goods API

var DevicePhone = &model.Device{Name: "phone", Price: 299}
var GoodsTV = &model.Device{Name: "tv", Price: 399}
var GoodsSwitch = &model.Device{Name: "switch", Price: 199}

func TestViewHandler_FindGoods(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name    string
		devices []*model.Device
		want    int
	}{
		{name: "returns one goods", devices: []*model.Device{DevicePhone}, want: 1},
		{name: "returns two goods", devices: []*model.Device{DevicePhone, GoodsTV}, want: 2},
		{name: "returns three goods", devices: []*model.Device{DevicePhone, GoodsTV, GoodsSwitch}, want: 3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(echo.GET, "/devices", nil)
			response := httptest.NewRecorder()
			store := &StubDeviceManager{}
			h := v1.NewViewHandler(store)
			insertDevices(t, store, tt.devices)

			c := e.NewContext(request, response)
			h.FindDevices(c)

			got := api.DecodeResponseV1(response.Body)
			fmt.Sprintf("%v", got)
			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, "", got.Message)
			assert.Equal(t, tt.want, len(got.Data.([]interface{})))
		})
	}
}

func TestViewHandler_CreateGoods(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name   string
		device *model.Device
		want   *model.Device
	}{
		{name: "create one goods", device: DevicePhone, want: DevicePhone},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, _ := json.Marshal(tt.device)
			request := httptest.NewRequest(echo.POST, "/goods", bytes.NewReader(jsonBytes))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			response := httptest.NewRecorder()
			store := &StubDeviceManager{}
			h := v1.NewViewHandler(store)

			c := e.NewContext(request, response)
			h.CreateDevice(c)

			got := api.DecodeResponseV1(response.Body)
			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, "", got.Message)
			reflect.DeepEqual(tt.want, got.Data)
		})
	}

	t.Run("unsupported media type", func(t *testing.T) {

	})
}

func TestViewHandler_EditGoods(t *testing.T) {
}

func TestBindUnmarshalParam(t *testing.T) {
	//store := &StubDeviceManager{}
	//h := v1.NewViewHandler(store)
	//e := app.NewEchoApp(h)
	e := echo.New()
	ts := model.Timestamp(time.Date(2016, 12, 6, 0, 0, 0, 0, time.UTC))
	jsonBytes, _ := json.Marshal(struct {
		Name  string `json:"name"`
		BuyAt string `json:"buyAt"`
	}{
		Name:  "tv",
		BuyAt: "2016-10-10",
	})
	fmt.Printf("%v", string(jsonBytes))
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBytes))
	//req := httptest.NewRequest(http.MethodPost, "/api/v1/device?buyAt=2016-12-06T00:00:00Z", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	result := struct {
		N string          `json:"name"`
		T model.Timestamp `json:"buyAt"`
	}{}
	err := c.Bind(&result)

	assert := assert.New(t)
	if assert.NoError(err) {
		//		assert.Equal( Timestamp(reflect.TypeOf(&Timestamp{}), time.Date(2016, 12, 6, 19, 9, 5, 0, time.UTC)), result.T)
		assert.Equal(ts, result.T)
		assert.Equal("tv", result.N)
	}
}

type StubDeviceManager struct {
	Devices []*model.Device
}

func (s *StubDeviceManager) GetOneDevice(id primitive.ObjectID) (*model.Device, error) {
	for _, item := range s.Devices {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, nil
}

func (s *StubDeviceManager) InsertOneDevice(item *model.Device) (*model.Device, error) {
	s.Devices = append(s.Devices, item)
	return item, nil
}

func (s *StubDeviceManager) GetAllDevices() []*model.Device {
	return s.Devices
}

func (s *StubDeviceManager) UpdateOneDevice(item *model.Device) (*model.Device, error) {
	for i, itemOld := range s.Devices {
		if itemOld.ID != item.ID {
			continue
		}
		s.Devices[i] = item
	}
	return item, nil
}

func insertDevices(t *testing.T, store *StubDeviceManager, devices []*model.Device) {
	t.Helper()

	for _, item := range devices {
		store.InsertOneDevice(item)
	}
}
