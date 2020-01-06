package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profile/api"
	"profile/auth"
	"profile/core"
	"time"
)

func (h *ViewHandler) Login(c echo.Context) error {
	data := c.(*core.CustomContext).GetBody()
	username := data["username"]
	password := data["password"]

	if username != "admin" || password != "123123" {
		return echo.ErrUnauthorized
	}

	t, err := auth.GenJWT("admin", []byte("secret-super-passwd"), time.Hour*6)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"code": api.CodeSuccess, "message": "登录成功", "token": t})
}

func (h *ViewHandler) UserInfo(c echo.Context) error {
	claims := c.(*core.CustomContext).GetUser()
	claims.Code = api.CodeSuccess
	return c.JSON(http.StatusOK, claims)
}

func (h *ViewHandler) Logout(c echo.Context) error {
	println("user logout")
	return c.JSON(http.StatusOK, echo.Map{})
}
