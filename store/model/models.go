package model

import "gorm.io/gorm"

// DatabaseConfig - структура для хранения конфигурации базы данных
type DatabaseConfig struct {
	gorm.Model
	Host     string
	Port     int
	Username string
	Password string
}

// Config - основная структура для хранения конфигурации
type Config struct {
	gorm.Model
	AppName    string
	Version    string
	Debug      bool
	DatabaseID uint
	Database   DatabaseConfig `gorm:"foreignKey:DatabaseID"`
}
