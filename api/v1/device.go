package v1

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"profile/api"
	"profile/model"
)

type IDeviceManager interface {
	InsertOneDevice(item *model.Device) (*model.Device, error)
	GetAllDevices() []*model.Device
	UpdateOneDevice(item *model.Device) (*model.Device, error)
	GetOneDevice(id primitive.ObjectID) (*model.Device, error)
	DeleteDeviceList(ids []primitive.ObjectID) (*mongo.DeleteResult, error)
}

// FindDevices GET to query devices records in db
// @Tags devices
// @Summary All devices
// @ID get-all-devices
// @Produce  json
// @Success 200 {object} model.Device[]
// @Header 200 {string} Token "qwerty"
// @Router /devices [get]
func (h *ViewHandler) FindDevices(c echo.Context) error {
	items := h.store.GetAllDevices()
	return c.JSON(http.StatusOK, api.ResponseV1(api.CodeSuccess, "", items))
}

// CreateDevice POST to create one new devices record in db
// @Tags device
// @Summary Create one new item
// @Description create new one
// @ID create-one-devices
// @Accept json
// @Produce json
// @Param device body model.Device true "add model.Device"
// @Header 200 {string} Authorization "Bearer qwerty"
// @Success 200 {object} model.Device
// @Router /device [post]
func (h *ViewHandler) CreateDevice(c echo.Context) error {
	item := model.NewDevice()
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	item, err := h.store.InsertOneDevice(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), item))
	}
	return c.JSON(http.StatusCreated, api.ResponseV1(api.CodeSuccess, "", item))
}

// @Summary EditDevice PUT to update devices in db
// @Tags devices
// @Description PUT method to update
// @ID edit-devices
// @Accept json
// @Produce json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} model.Device
// @Router /devices [put]
func (h *ViewHandler) EditDevice(c echo.Context) error {
	item := new(model.Device)
	if err := c.Bind(item); err != nil {
		return err
	}
	itemID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	itemModel, err := h.store.GetOneDevice(itemID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), itemModel))
	}

	item, err = h.store.UpdateOneDevice(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), item))
	}
	return c.JSON(http.StatusOK, api.ResponseV1(api.CodeSuccess, "", item))
}

// @Summary DeleteDevice DELETE to delete devices in db
// @Tags devices
// @Description PUT method to update
// @ID delete-devices
// @Accept json
// @Produce json
// @Success 200 {object} model.Device
// @Router /devices [delete]
func (h *ViewHandler) DeleteDevice(c echo.Context) error {
	var data struct {
		IDs []string `json:"ids"`
	}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), 0))
	}

	var objectIDs []primitive.ObjectID
	for _, id := range data.IDs {
		itemID, _ := primitive.ObjectIDFromHex(id)
		objectIDs = append(objectIDs, itemID)
	}
	delRet, err := h.store.DeleteDeviceList(objectIDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(api.CodeSuccess, err.Error(), 0))
	}

	return c.JSON(http.StatusOK, api.ResponseV1(api.CodeSuccess, "删除成功", delRet.DeletedCount))
}
