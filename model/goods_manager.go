package model

import "github.com/jinzhu/gorm"

type GoodsManger struct {
	db *gorm.DB
}

func NewGoodsManager(db *gorm.DB) *GoodsManger {
	return &GoodsManger{
		db: db,
	}
}

func (m *GoodsManger) InsertOneGoods(item Goods) Goods {
	if m.db.NewRecord(item) {
		m.db.Create(&item)
	}
	return item
}

func (m *GoodsManger) GetAllGoods() []Goods {
	var goods []Goods
	m.db.Find(&goods)
	return goods
}
