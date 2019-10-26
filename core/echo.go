package core

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"profile/auth"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) GetBody() echo.Map {
	m := echo.Map{}
	if c.Request().Body != nil {
		c.Bind(&m)
	}
	return m
}

func (c *CustomContext) GetUser() *auth.JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*auth.JwtCustomClaims)
}
