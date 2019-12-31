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
	Name              string             `json:"name" bson:"name"`
	Description       string             `bson:"description"`
	Price             uint               `bson:"price"`
	Category          uint               `bson:"category"`
	Images            []Image            `bson:"images"`
	BuyAt             Timestamp          `json:"buyAt" bson:"buyAt"`
	ExpiredAt         Timestamp          `bson:"expiredAt"`
}

type Image struct {
	BaseModelWithTime `bson:",inline"`
	Path              string `bson:"path"`
}
