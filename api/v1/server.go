package v1

import (
	"net/http"
	"profile/model"

	"github.com/labstack/echo"
)

type Store interface {
	GoodsManager
}

type GoodsManager interface {
	GetAllGoods() []model.Goods
}

type ViewHandler struct {
	store Store
}

func NewViewHandler(store Store) *ViewHandler {
	return &ViewHandler{
		store: store,
	}
}

func (h *ViewHandler) Profiles(c echo.Context) error {
	username := c.Param("username")
	println(username)
	return c.String(http.StatusOK, GetUserProfile(username))
}

func GetUserProfile(username string) string {
	if username == "Peter" {
		return "Peter's Profile"

	}
	if username == "Chris" {
		return "Chris's Profile"
	}

	return ""
}
