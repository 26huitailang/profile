package v1

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *ViewHandler) Goods(c echo.Context) error {
	items := h.store.GetAllGoods()
	return c.JSON(http.StatusOK, items)
}
