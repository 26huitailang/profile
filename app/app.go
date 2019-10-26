package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	v1 "profile/api/v1"
	"profile/auth"
	"profile/core"
)

func NewEchoApp(h *v1.ViewHandler) *echo.Echo {

	e := echo.New()
	ConfigCustomContext(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/profiles/:username", h.Profiles)
	e.POST("/login", h.Login)

	apiRoute := e.Group("/api")
	// basic auth
	apiRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret-super-passwd"),
		Claims:     &auth.JwtCustomClaims{},
	}))
	apiRoute.GET("/check-login", h.CheckLogin)

	apiV1 := apiRoute.Group("/v1")
	{
		apiV1.GET("/goods", h.FindGoods)
		apiV1.POST("/goods", h.CreateGoods)
		apiV1.PUT("/goods/:id", h.EditGoods)
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
