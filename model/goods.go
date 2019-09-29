package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GoodsCategory uint

const (
	Others GoodsCategory = iota // 0
	Electronics
)

type Goods struct {
	gorm.Model
	Name        string
	Description string
	Price       uint
	Category    uint
	Images      []GoodsImage
	Profile     GoddsProfile
	ProfileID   uint
}

type GoodsImage struct {
	gorm.Model
	GoodsID uint
	Path    string
}

type GoddsProfile struct {
	gorm.Model
	BuyAt            time.Time
	ExpireAt         time.Time
	DepreciationRate float32
}
