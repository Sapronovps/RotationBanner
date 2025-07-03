package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConf
	Logger LoggerConf
	DB     DBConf
}

type ServerConf struct {
	Host        string
	Port        int
	AddressGrpc string
	IsHTTP      bool
}

type LoggerConf struct {
	File  string
	Level string
}

type DBConf struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	InMemory bool
}

func NewConfig(configPath string) Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetEnvPrefix("db")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("decode into struct: %w", err))
	}

	return c
}
