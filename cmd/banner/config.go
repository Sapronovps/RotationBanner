package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConf
	Logger LoggerConf
	DB     DBConf
	Kafka  KafkaConf
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

type KafkaConf struct {
	Brokers     string
	RetryMax    int
	EventsTopic string
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

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Задаем переменные окружения вручную, чтобы указать маппинг
	viper.BindEnv("server.host")
	viper.BindEnv("server.port")
	viper.BindEnv("server.addressgrpc")
	viper.BindEnv("server.ishttp")

	viper.BindEnv("logger.file")
	viper.BindEnv("logger.level")

	viper.BindEnv("db.host")
	viper.BindEnv("db.port")
	viper.BindEnv("db.username")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")
	viper.BindEnv("db.inmemory")

	viper.BindEnv("kafka.brokers")
	viper.BindEnv("kafka.retrymax")
	viper.BindEnv("kafka.eventstopic")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения конфигурации: %w", err))
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("decode into struct: %w", err))
	}

	return c
}
