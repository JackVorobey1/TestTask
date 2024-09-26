package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter struct {
	DB *gorm.DB
}

func NewAdapter(cfg *Config) (*PostgresAdapter, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", cfg.Host, cfg.Username, cfg.Password, cfg.DBname, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Возвращаем адаптер с подключением к базе данных
	return &PostgresAdapter{
		DB: db,
	}, nil
}
