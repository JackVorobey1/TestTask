package model

//?????????????????//
import (
	"TestTask/store/adapters/postgresql"
	"gorm.io/gorm"
)

// Store структура, содержащая подключение к базе данных
type Store struct {
	DB *gorm.DB
}

// CreateStore создаёт и возвращает объект Store с подключением к базе данных
func CreateStore(config postgresql.Config) (*Store, error) {
	// Подключаемся к базе данных PostgreSQL через адаптер
	db, err := postgresql.ConnectPsQl(&config)
	if err != nil {
		return nil, err
	}

	// Возвращаем объект Store с подключением GORM
	return &Store{
		DB: db,
	}, nil
}
