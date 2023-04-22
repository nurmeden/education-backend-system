package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)

	// Указываем, что используем переменные окружения для настройки
	v.SetEnvPrefix("MY_MICROSERVICE")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Указываем значения по умолчанию, если они не указаны
	v.SetDefault("port", "8080")
	v.SetDefault("database_url", "postgres://user:password@localhost/dbname")

	// Создаем экземпляр структуры Config и привязываем значения конфигурации
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return cfg, nil
}
