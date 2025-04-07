package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// Config структура для хранения конфигурации
type Config struct {
	APIKey          string
	APIURL          string
	AvailableModels []string
}

// LoadConfig загружает конфигурацию из .env файла
func LoadConfig() (*Config, error) {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла, используем переменные окружения")
	}

	// Получаем значения из .env
	config := &Config{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		APIURL: os.Getenv("OPENAI_API_URL"),
	}

	// Загружаем доступные модели из переменной окружения
	availableModels := os.Getenv("OPENAI_AVAILABLE_MODELS")
	if availableModels == "" {
		return nil, log.Output(2, "Переменная окружения OPENAI_AVAILABLE_MODELS не найдена")
	}

	// Разделяем строку моделей на срез
	config.AvailableModels = strings.Split(availableModels, ",")

	return config, nil
}
