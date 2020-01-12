package model_test

import (
	"github.com/magiconair/properties/assert"
	"profile/auth"
	"profile/model"
	"testing"
)

func TestUserManager_InsertOneUser(t *testing.T) {
	helperConfigInit()
	t.Run("val unique", func(t *testing.T) {
		user1 := model.NewUser("user1", "user1@mail.com", auth.EncryptPassword("password"))
		user2 := model.NewUser("user1", "user2@mail.com", auth.EncryptPassword("password"))
		client := initMongoClient(t)
		um := model.NewUserManager(client)
		defer helperDropCollection(um)

		_, err := um.InsertOneUser(user1)
		checkError(t, err)
		_, err = um.InsertOneUser(user2)
		if err == nil {
			t.Fatal("expect duplicate key error")
		}
	})
	t.Run("email unique", func(t *testing.T) {
		user1 := model.NewUser("user1", "user1@mail.com", auth.EncryptPassword("password"))
		user2 := model.NewUser("user2", "user1@mail.com", auth.EncryptPassword("password"))
		client := initMongoClient(t)
		um := model.NewUserManager(client)
		defer helperDropCollection(um)

		_, err := um.InsertOneUser(user1)
		checkError(t, err)
		_, err = um.InsertOneUser(user2)
		if err == nil {
			t.Fatal("expect duplicate key error")
		}
	})
}

func TestUserManager_FindOneUser(t *testing.T) {
	helperConfigInit()
	user1 := model.NewUser("user1", "user1@mail.com", auth.EncryptPassword("password"))
	user2 := model.NewUser("user2", "user2@mail.com", auth.EncryptPassword("password"))

	type args struct {
		key   string
		val   string
		users []*model.User
	}
	cases := []struct {
		name string
		args args
		want string
	}{
		{name: "find one by val ok", args: args{key: "username", val: "user1", users: []*model.User{user1, user2}}, want: user1.Username},
		{name: "find one by email ok", args: args{key: "email", val: "user1@mail.com", users: []*model.User{user1, user2}}, want: user1.Username},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := initMongoClient(t)
			um := model.NewUserManager(client)
			defer helperDropCollection(um)

			for _, item := range tt.args.users {
				_, _ = um.InsertOneUser(item)
			}
			got, err := um.FindOneUser(tt.args.key, tt.args.val)
			checkError(t, err)
			assert.Equal(t, tt.want, got.Username)
		})
	}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}

}
