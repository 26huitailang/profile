package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profile/api"
	"profile/model"
	"strconv"
)

type IDeviceManager interface {
	InsertOneDevice(item *model.Device) (*model.Device, error)
	GetAllDevices() []*model.Device
	UpdateOneDevice(item *model.Device) (*model.Device, error)
	GetOneDevice(id uint) (*model.Device, error)
}

// FindDevices GET to query goods records in db
// @Tags goods
// @Summary All goods
// @ID get-all-goods
// @Produce  json
// @Success 200 {object} model.Goods[]
// @Header 200 {string} Token "qwerty"
// @Router /goods [get]
func (h *ViewHandler) FindDevices(c echo.Context) error {
	items := h.store.GetAllDevices()
	return c.JSON(http.StatusOK, api.ResponseV1(api.CodeSuccess, "", items))
}

// CreateDevice POST to create one new goods record in db
// @Tags goods
// @Summary Create one new item
// @Description create new one
// @ID create-one-goods
// @Accept json
// @Produce json
// @Param name body model.Goods true "add model.Goods"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} model.Goods
// @Router /goods [post]
func (h *ViewHandler) CreateDevice(c echo.Context) error {
	item := new(model.Device)
	if err := c.Bind(item); err != nil {
		return err
	}
	item, err := h.store.InsertOneDevice(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), item))
	}
	return c.JSON(http.StatusCreated, api.ResponseV1(api.CodeSuccess, "", item))
}

// @Summary EditGoods PUT to update goods in db
// @Summary EditGoods PUT to update goods
// @Tags goods
// @Description PUT method to update
// @ID edit-goods
// @Accept json
// @Produce json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} model.Goods
// @Router /goods [put]
func (h *ViewHandler) EditGoods(c echo.Context) error {
	item := new(model.Device)
	if err := c.Bind(item); err != nil {
		return err
	}
	itemID, _ := strconv.Atoi(c.Param("id"))
	itemModel, err := h.store.GetOneDevice(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), itemModel))
	}

	item, err = h.store.UpdateOneDevice(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), item))
	}
	return c.JSON(http.StatusOK, api.ResponseV1(api.CodeSuccess, "", item))
}
