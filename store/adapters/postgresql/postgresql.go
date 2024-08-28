package postgresql

import (
	"TestTask/store/model"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsn := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Cannot start %s", err)
	}

	db.AutoMigrate(&model.Config{})
	return db
}

// ?????//
func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		sleep := 1 * time.Second
		maxAttempts := 5
		attempts := 0

		for dbase == nil && attempts < maxAttempts {
			fmt.Printf("База данных недоступна. Подождите %v сек.\n", sleep.Seconds())
			time.Sleep(sleep)
			sleep *= 2 // удваиваем время ожидания
			dbase = Init()
			attempts++
		}

		if dbase == nil {
			log.Fatal().Msg("Не удалось установить соединение с базой данных после нескольких попыток")
		}
	}
	return dbase
}
