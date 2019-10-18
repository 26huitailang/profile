package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"net/http"
	v1 "profile/api/v1"
	"profile/database"
	"profile/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	_ "profile/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
func main() {
	db, closeDB := database.NewDB("prod.db")
	defer closeDB()

	// 自动迁移模式
	db.AutoMigrate(&model.Goods{}, &model.GoodsImage{})

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	store := model.NewGoodsManager(db)
	h := v1.NewViewHandler(store)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/profiles/:username", h.Profiles)

	apiRoute := e.Group("/api")

	apiV1 := apiRoute.Group("/v1")
	apiV1.GET("/goods", h.FindGoods)
	apiV1.POST("/goods", h.CreateGoods)
	apiV1.PUT("/goods/:id", h.EditGoods)

	e.Logger.Fatal(e.Start(":5000"))
}
