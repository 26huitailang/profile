package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Others              uint = iota // 0
	ElectronicEquipment             // 1
	HouseholdAppliances             // 2
)

type BaseModelWithTime struct {
	CreatedAt Timestamp `bson:"createdAt"`
	UpdatedAt Timestamp `bson:"updatedAt"`
}

type Device struct {
	BaseModelWithTime `bson:",inline"`
	ID                primitive.ObjectID `bson:"_id"`
	Name              string             `json:"name" bson:"name" validate:"required"`
	Description       string             `bson:"description"`
	Price             uint               `bson:"price" example:"9.9"`
	Category          uint               `bson:"category" enums:"0,1,2" example:"1" validate:"required"`
	Images            []Image            `bson:"images"`
	BuyAt             Timestamp          `json:"buyAt" bson:"buyAt"`
	ExpiredAt         Timestamp          `bson:"expiredAt"`
}

type Image struct {
	BaseModelWithTime `bson:",inline"`
	Path              string `bson:"path"`
}
