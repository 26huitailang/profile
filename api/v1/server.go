package v1

import (
	"net/http"

	"github.com/labstack/echo"
)

type Store interface{}

type ViewHandler struct {
	store Store
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
