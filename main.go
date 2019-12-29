package main

import (
	"log"
	v1 "profile/api/v1"
	"profile/app"
	"profile/config"
	"profile/database"
	"profile/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	// conifg
	client, err := database.NewMongo("root", "123123", database.MongoHost, database.MongoDB)
	if err != nil {
		log.Fatalf("DB connect error: %s", err)
	}
	// 自动迁移模式
	//db.AutoMigrate(&model.Goods{}, &model.GoodsImage{})

	store := model.NewDeviceManager(client)
	h := v1.NewViewHandler(store)

	e := app.NewEchoApp(h)
	e.Logger.Fatal(e.Start(config.GetString("server.port")))
}
