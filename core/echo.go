package core

import "github.com/labstack/echo/v4"

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
