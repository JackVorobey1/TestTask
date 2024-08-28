package internal

import (
	"TestTask/store/model"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
)

// Load - функция для чтения и анмаршалинга файла config.yaml
func Load() {
	var config model.Config
	file, err := os.ReadFile("templates/config.yaml") // читаем файл
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка при чтении файла config.yaml")
	}
	// Анмаршалинг YAML данных в структуру Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка при анмаршалинге файла config.yaml")
	}

	// Вывод данных на экран
	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Printf("Debug: %v\n", config.Debug)
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)

	log.Info().Msg("Конфигурация загружена успешно")
}
