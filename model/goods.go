package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GoodsCategory uint

const (
	Others GoodsCategory = iota // 0
	Electronics
)

type CustomModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

type Goods struct {
	CustomModel
	Name             string       `json:"name" gorm:"not null;unique" example:"item name"`
	Description      string       `json:"description" example:"description"`
	Price            uint         `json:"price" example:"9"`
	Category         uint         `json:"category"`
	Images           []GoodsImage `json:"images"`
	BuyAt            time.Time    `json:"buyAt"`
	ExpireAt         time.Time    `json:"expireAt"`
	DepreciationRate float32      `json:"depreciationRate"` // percent per day
}

type GoodsImage struct {
	CustomModel
	GoodsID uint   `json:"goodsID"`
	Path    string `json:"path"`
}
