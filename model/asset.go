package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Others              uint = iota // 0
	ElectronicEquipment             // 1
	HouseholdAppliances             // 2
)

type BaseModelWithTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Device struct {
	BaseModelWithTime
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Description string
	Price       uint
	Category    uint
	Images      []Image
	BuyAt       time.Time
	ExpiredAt   time.Time
}

type Image struct {
	BaseModelWithTime
	Path string
}
