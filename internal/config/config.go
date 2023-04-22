package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"port"`
	DatabaseUrl string `mapstructure:"database_url"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
