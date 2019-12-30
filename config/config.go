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
}

var Cfg Config

func init() {
	// default
	viper.SetDefault("level", LevelDevelop)

	// flags
	parseFlag()

	// config file
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	overrideConfig(viper.GetString("level"))

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

func overrideConfig(level string) {
	switch level {
	case LevelTest:
		viper.Set("db.name", DBNameTest)
	case LevelDebug:
		viper.Set("db.name", DBNameDev)
	case LevelDevelop:
		viper.Set("db.name", DBNameDev)
	case LevelProduct:
		viper.Set("db.name", DBNameProd)
	default:
		err := fmt.Errorf("no this level %d", level)
		fmt.Print(err)
		os.Exit(-1)
	}
}

func parseFlag() {
	pflag.String("level", LevelDevelop, "env level, default develop, [test/debug/develop/product]")
	pflag.String("server.port", Port, "server port, default :5000")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
