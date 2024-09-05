package postgresql

import (
	"fmt"
	"gorm.io/gorm"
)

func ConnectPsQl(config *Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s username=%s password=%s port=%d", config.Host, config.Username, config.Password, config.Port)
	return
}
