package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	LevelTest    = "test"
	LevelDebug   = "debug"
	LevelDevelop = "develop"
	LevelProduct = "product"
)

const (
	DBNameTest = "test.db"
	DBNameDev  = "dev.db"
	DBNameProd = "prod.db"
	Port       = ":5000"
)

type DB struct {
	Name string `mapstructure:"name"`
}
type Mongo struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DB       string `mapstructure:"db"`
}
type Server struct {
	Port string `mapstructure:"port"`
}
type Config struct {
	Level  string
	DB     `mapstructure:"db"`
	Server `mapstructure:"server"`
	Mongo  `mapstructure:"mongo"`
}

var Cfg Config

func InitConfig() {
	// default
	configType := "yaml"
	defaultPath1 := "./config"
	defaultPath2 := "."

	v := viper.New()
	// config file
	v.SetConfigName("default") // name of config file (without extension)
	v.SetConfigType(configType)
	v.AddConfigPath(defaultPath1) // optionally look for config in the working directory
	v.AddConfigPath(defaultPath2) // optionally look for config in the working directory

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	configs := v.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("GO_ENV")

	if env != "" {
		viper.SetConfigName(env)
		viper.AddConfigPath(defaultPath1)
		viper.AddConfigPath(defaultPath2)
		viper.SetConfigType(configType)
		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	// flags
	//parseFlag()

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Printf("%+v\n", Cfg)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func Sub(key string) *viper.Viper {
	return viper.Sub(key)
}

func parseFlag() {
	pflag.String("server.port", Port, "server port, default :5000")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
