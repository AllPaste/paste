package config

import (
	"log"

	"github.com/spf13/viper"
)

var c Config

type Config struct {
	DB     *DB     `mapstructure:"DB"`
	Log    *Log    `mapstructure:"Log"`
	Server *Server `mapstructure:"Server"`
}

type Log struct {
	Level string `mapstructure:"Level"`
}

type DB struct {
	Driver string `mapstructure:"Driver"`
	DSN    string `mapstructure:"DSN"`
}

type Server struct {
	Listen string `mapstructure:"Listen"`
}

func NewConfig(path string) *Config {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed: %v", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("unmarshal config file failed: %v", err)
	}

	return &c
}
