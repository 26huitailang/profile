package config_test

import (
	"profile/config"
	"testing"
)

func TestConfig_Mongo(t *testing.T) {
	tests := []struct {
		name        string
		configLevel string
		configMongo config.Mongo
	}{
		{"level develop", "develop", config.Mongo{DB: "develop", Host: "127.0.0.1:27017", Username: "develop", Password: "develop"}},
		{"level test", "test", config.Mongo{DB: "develop", Host: "127.0.0.1:27017", Username: "test", Password: "test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
	config.Cfg.Level
	config.Mongo{
		Username: "",
		Password: "",
		Host:     "",
		DB:       "",
	}
}
