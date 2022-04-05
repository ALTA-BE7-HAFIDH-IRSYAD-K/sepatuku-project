package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port      int    `yaml:"port"`
	SecretJWT string `yaml:"secretjwt"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.SecretJWT = "S3CR3T"
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "db-name"
	defaultConfig.Database.Address = "127.0.0.1"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "username"
	defaultConfig.Database.Password = "password"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("failed to open file")
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract external config, use default value")
		return &defaultConfig
	}
	return &finalConfig
}
