package config

import (
	"errors"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	key := os.Getenv("JSONBIN_KEY")
	if key == "" {
		return nil, errors.New("переменная окружения JSONBIN_KEY не установлена")
	}
	return &Config{Key: key}, nil
}