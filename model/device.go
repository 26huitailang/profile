package model

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

const (
	CategoryOthers              uint = iota // 0
	CategoryElectronicEquipment             // 1
	CategoryHouseholdAppliances             // 2
)

const (
	DeviceStatusUsing string = "using" // 使用中
	DeviceStatusStack string = "stack" // 收起
)

type BaseModelWithTime struct {
	CreatedAt Timestamp `json:"createdAt" bson:"createdAt"`
	UpdatedAt Timestamp `json:"updatedAt" bson:"updatedAt"`
}

func NewDevice() *Device {
	return &Device{
		BaseModelWithTime: BaseModelWithTime{CreatedAt: Timestamp(time.Now()), UpdatedAt: Timestamp(time.Now())},
		ID:                primitive.NewObjectID(),
		Name:              "",
		Description:       "",
		Price:             0.0,
		Category:          0,
		Images:            []Image{},
		BuyAt:             Timestamp{},
		ExpiredAt:         Timestamp{},
		Status:            DeviceStatusUsing,
	}
}

type Device struct {
	BaseModelWithTime `bson:",inline"`
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Name              string             `json:"name" bson:"name" validate:"required"`
	Description       string             `json:"description" bson:"description"`
	Price             PriceType          `json:"price" bson:"price" example:"9.9"`
	Category          uint               `json:"category" bson:"category" enums:"0,1,2" example:"1" validate:"required"`
	Images            []Image            `json:"images" bson:"images"`
	BuyAt             Timestamp          `json:"buyAt" bson:"buyAt"`
	ExpiredAt         Timestamp          `json:"expiredAt" bson:"expiredAt"`
	Status            string             `json:"status" bson:"status"`
}

type Image struct {
	BaseModelWithTime `bson:",inline"`
	Path              string `json:"path" bson:"path"`
}

type PriceType float64

func (p *PriceType) UnmarshalJSON(src []byte) error {
	src = bytes.Trim(src, "\"")
	price, err := strconv.ParseFloat(string(src), 64)
	if err != nil {
		return err
	}
	*p = PriceType(price)
	return nil
}
