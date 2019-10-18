package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profile/api"
	"profile/model"
	"strconv"
)

type GoodsManager interface {
	InsertOneGoods(item *model.Goods) (*model.Goods, error)
	GetAllGoods() []model.Goods
	UpdateOneGoods(item *model.Goods) (*model.Goods, error)
	GetOneGoods(id uint) (*model.Goods, error)
}

// FindGoods GET to query goods records in db
// @Tags goods
// @Summary All goods
// @ID get-all-goods
// @Produce  json
// @Success 200 {object} model.Goods[]
// @Header 200 {string} Token "qwerty"
// @Router /goods [get]
func (h *ViewHandler) FindGoods(c echo.Context) error {
	items := h.store.GetAllGoods()
	return c.JSON(http.StatusOK, api.ResponseV1("", items))
}

// CreateGoods POST to create one new goods record in db
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
func (h *ViewHandler) CreateGoods(c echo.Context) error {
	item := new(model.Goods)
	if err := c.Bind(item); err != nil {
		return err
	}
	item, err := h.store.InsertOneGoods(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(err.Error(), item))
	}
	return c.JSON(http.StatusCreated, api.ResponseV1("", item))
}

// EditGoods PUT to update goods in db
func (h *ViewHandler) EditGoods(c echo.Context) error {
	item := new(model.Goods)
	if err := c.Bind(item); err != nil {
		return err
	}
	itemID, _ := strconv.Atoi(c.Param("id"))
	itemModel, err := h.store.GetOneGoods(uint(itemID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(err.Error(), itemModel))
	}

	item, err = h.store.UpdateOneGoods(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, api.ResponseV1(err.Error(), item))
	}
	return c.JSON(http.StatusOK, api.ResponseV1("", item))
}
