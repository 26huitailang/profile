package database

import (
	"github.com/jinzhu/gorm"
)

func NewDB(filename string) (*gorm.DB, func()) {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic("连接数据库失败")
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	return db, func() {
		db.Close()
	}
}
