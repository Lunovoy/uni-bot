package config

import (
	"os"
)

// Структура для хранения конфигурации
type Config struct {
	TelegramToken string
	OpenAIToken   string
}

// Загрузка конфигурации из файла .env
func NewConfig() *Config {
	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		OpenAIToken:   os.Getenv("OPENAI_TOKEN"),
	}
}
