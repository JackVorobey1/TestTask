package adapters

import "gorm.io/gorm"

// PostgresStore - структура для хранения соединения с базой данных
type PostgresStore struct {
	db *gorm.DB
}

//????????????//
