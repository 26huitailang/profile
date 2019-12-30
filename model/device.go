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

type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	layout := time.RFC3339
	ts, err := time.Parse(layout, src)
	if err == nil {

		*t = Timestamp(ts)
		return nil
	}
	layout = "2006-01-02"
	ts, err = time.Parse(layout, src)
	*t = Timestamp(ts)
	return err
}

type BaseModelWithTime struct {
	CreatedAt Timestamp `bson:"createdAt"`
	UpdatedAt Timestamp `bson:"updatedAt"`
}

type Device struct {
	BaseModelWithTime
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `bson:"description"`
	Price       uint               `bson:"price"`
	Category    uint               `bson:"category"`
	Images      []Image            `bson:"images"`
	BuyAt       Timestamp          `json:"buyAt" bson:"buyAt"`
	ExpiredAt   Timestamp          `bson:"expiredAt"`
}

type Image struct {
	BaseModelWithTime
	Path string `bson:"path"`
}
