package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	v1 "profile/api/v1"
	"profile/auth"
	"profile/core"
)

func NewEchoApp(h *v1.ViewHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())

	// todo: debug mode
	e.Debug = true
	ConfigCustomContext(e)

	if e.Debug {
		e.Use(middleware.BodyDump(func(context echo.Context, req []byte, res []byte) {
			fmt.Fprintf(os.Stderr, "Request: %v\n", string(req))
		}))
		e.Use(middleware.Logger())
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Token"},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/profiles/:username", h.Profiles)
	e.POST("/user/login", h.Login)
	e.POST("/user/logout", h.Logout)

	apiAuthRoute := e.Group("")

	// jwt auth
	apiAuthRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret-super-passwd"),
		Claims:      &auth.JwtCustomClaims{},
		TokenLookup: "header:Authorization",
	}))
	apiAuthRoute.GET("/user/info", h.UserInfo)

	apiV1 := apiAuthRoute.Group("/api/v1")
	{
		apiV1.GET("/devices", h.FindDevices)
		apiV1.POST("/device", h.CreateDevice)
		apiV1.PUT("/devices/:id", h.EditDevice)
	}
	return e
}

func ConfigCustomContext(e *echo.Echo) {
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			cc := &core.CustomContext{context}
			return handlerFunc(cc)
		}
	})
}
