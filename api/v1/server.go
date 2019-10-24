package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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
