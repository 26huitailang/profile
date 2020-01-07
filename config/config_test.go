package config_test

import (
	"github.com/stretchr/testify/assert"
	"os"
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
		{"level test", "test", config.Mongo{DB: "test", Host: "127.0.0.1:27017", Username: "test", Password: "test"}},
	}

	for _, tt := range tests {
		_ = os.Setenv("GO_ENV", tt.configLevel)
		config.InitConfig()
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v", config.Cfg)
			assert.Equal(t, config.Cfg.Mongo.DB, tt.configMongo.DB)
			assert.Equal(t, config.Cfg.Mongo.Username, tt.configMongo.Username)
			assert.Equal(t, config.Cfg.Mongo.Password, tt.configMongo.Password)
		})
	}
}
