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

func (m *GoodsManger) InsertOneGoods(item *Goods) (*Goods, error) {
	if err := m.db.Create(item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (m *GoodsManger) GetAllGoods() []Goods {
	var goods []Goods
	m.db.Find(&goods)
	return goods
}

func (m *GoodsManger) UpdateOneGoods(item *Goods) (*Goods, error) {
	err := m.db.Save(item).Error
	return item, err
}

func (m *GoodsManger) GetOneGoods(id uint) (*Goods, error) {
	item := new(Goods)
	err := m.db.First(item, id).Error

	return item, err
}
