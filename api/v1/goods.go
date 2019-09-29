package v1

import (
	"net/http"
	"profile/model"

	"github.com/labstack/echo"
)

type GoodsManager interface {
	InsertOneGoods(item model.Goods) model.Goods
	GetAllGoods() []model.Goods
}

func (h *ViewHandler) GetGoods(c echo.Context) error {
	items := h.store.GetAllGoods()
	return c.JSON(http.StatusOK, items)
}

func (h *ViewHandler) PostGoods(c echo.Context) error {
	item := c.Bind()
	item := h.store.InsertOneGoods()
	return c.JSON(http.StatusOK, items)
}
