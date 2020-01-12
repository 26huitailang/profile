package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewUser(username, email, password string) *User {
	return &User{
		BaseModelWithTime: BaseModelWithTime{CreatedAt: Timestamp(time.Now()), UpdatedAt: Timestamp(time.Now())},
		ID:                primitive.NewObjectID(),
		Username:          username,
		Email:             email,
		Password:          password,
		LastLoginAt:       Timestamp{},
		IsActive:          true,
	}
}

type User struct {
	BaseModelWithTime `bson:",inline"`
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Username          string             `json:"username" bson:"username" validate:"required"`
	Password          string             `json:"password" bson:"password" validate:"required"`
	Email             string             `json:"email" bson:"email"`
	LastLoginAt       Timestamp          `json:"lastLoginAt" bson:"lastLoginAt"`
	IsActive          bool               `json:"isActive" bson:"isActive" default:"true"`
}
