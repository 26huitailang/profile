package main

import (
	v1 "profile/api/v1"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	h := &v1.ViewHandler{}

	e.GET("/profiles/:username", h.Profiles)

	e.Logger.Fatal(e.Start(":1323"))
}
