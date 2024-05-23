package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error loading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Error on unmarshal of config file, %s", err)
		return
	}

	return
}
