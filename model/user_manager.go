package model

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"profile/config"
)

type UserManager struct {
	BaseManager
}

func NewUserManager(client *mongo.Client) *UserManager {
	um := &UserManager{
		BaseManager: BaseManager{
			collection: client.Database(config.Cfg.Mongo.DB).Collection("user"),
		},
	}
	_, err := CreateIndexes(um.collection, "val")
	if err != nil {
		log.Fatal(err)
	}
	_, err = CreateIndexes(um.collection, "email")
	if err != nil {
		log.Fatal(err)
	}
	return um
}

func (m *UserManager) InsertOneUser(item *User) (ret *mongo.InsertOneResult, err error) {
	ret, err = m.InsertOne(item)
	return
}

// FindOneUser
// key: val/email
func (m *UserManager) FindOneUser(key, val string) (*User, error) {
	var user User
	ret := m.FindOne(bson.D{{key, val}})
	err := ret.Decode(&user)
	return &user, err
}
