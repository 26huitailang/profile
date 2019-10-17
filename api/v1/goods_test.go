package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"profile/api"
	v1 "profile/api/v1"
	"profile/model"
	"reflect"
	"testing"
)

// todo: unittest for Goods API

var GoodsPhone = model.Goods{Name: "phone", Price: 299}
var GoodsTV = model.Goods{Name: "tv", Price: 399}
var GoodsSwitch = model.Goods{Name: "switch", Price: 199}

func TestViewHandler_FindGoods(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name  string
		goods []model.Goods
		want  int
	}{
		{name: "returns one goods", goods: []model.Goods{GoodsPhone}, want: 1},
		{name: "returns two goods", goods: []model.Goods{GoodsPhone, GoodsTV}, want: 2},
		{name: "returns three goods", goods: []model.Goods{GoodsPhone, GoodsTV, GoodsSwitch}, want: 3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(echo.GET, "/goods", nil)
			response := httptest.NewRecorder()
			store := &StubGoodsManager{}
			h := v1.NewViewHandler(store)
			insertGoods(t, store, tt.goods)

			c := e.NewContext(request, response)
			h.FindGoods(c)

			got := api.DecodeResponseV1(response.Body)
			fmt.Sprintf("%v", got)
			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, "", got.Info)
			assert.Equal(t, tt.want, len(got.Data.([]interface{})))
		})
	}
}

func TestViewHandler_CreateGoods(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name  string
		goods model.Goods
		want  model.Goods
	}{
		{name: "create one goods", goods: GoodsPhone, want: GoodsPhone},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, _ := json.Marshal(tt.goods)
			request := httptest.NewRequest(echo.POST, "/goods", bytes.NewReader(jsonBytes))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			response := httptest.NewRecorder()
			store := &StubGoodsManager{}
			h := v1.NewViewHandler(store)

			c := e.NewContext(request, response)
			h.CreateGoods(c)

			got := api.DecodeResponseV1(response.Body)
			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, "", got.Info)
			reflect.DeepEqual(tt.want, got.Data)
		})
	}

	t.Run("unsupported media type", func(t *testing.T) {

	})
}

func TestViewHandler_EditGoods(t *testing.T) {
}

type StubGoodsManager struct {
	Goods []model.Goods
}

func (s *StubGoodsManager) GetOneGoods(id uint) (*model.Goods, error) {
	for _, item := range s.Goods {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, nil
}

func (s *StubGoodsManager) InsertOneGoods(item *model.Goods) (*model.Goods, error) {
	s.Goods = append(s.Goods, *item)
	return item, nil
}

func (s *StubGoodsManager) GetAllGoods() []model.Goods {
	return s.Goods
}

func (s *StubGoodsManager) UpdateOneGoods(item *model.Goods) (*model.Goods, error) {
	for i, itemOld := range s.Goods {
		if itemOld.ID != item.ID {
			continue
		}
		s.Goods[i] = *item
	}
	return item, nil
}

func insertGoods(t *testing.T, store *StubGoodsManager, goods []model.Goods) {
	t.Helper()

	for _, item := range goods {
		store.InsertOneGoods(&item)
	}
}
