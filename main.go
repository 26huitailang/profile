package main

import (
	v1 "profile/api/v1"
	"profile/database"
	"profile/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {
	db, closeDB := database.NewDB("prod.db")
	defer closeDB()

	// 自动迁移模式
	db.AutoMigrate(&model.Goods{}, &model.GoddsProfile{}, &model.GoodsImage{})

	e := echo.New()
	store := model.NewGoodsManager(db)
	h := v1.NewViewHandler(store)

	e.GET("/profiles/:username", h.Profiles)
	e.GET("/goods", h.Goods)

	e.Logger.Fatal(e.Start(":5000"))
}
