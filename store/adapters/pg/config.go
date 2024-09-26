package pg

import (
	"errors"
	"strings"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	DBname   string `yaml:"dbname"`
	Password string `yaml:"password"`
}

// метод Validdate структуры Config
func (c *Config) Validdate() (err error) {

	if c == nil {
		return errors.New("Empty adapter config")
	}

	//удаляет пробелы с начала и конца строки
	if strings.TrimSpace(c.Host) == "" {
		err = errors.New("Invalid Version ")
		return
	}

	if c.Port <= 0 {
		err = errors.New("Invalid Port")
		return
	}

	if strings.TrimSpace(c.Username) == "" {
		err = errors.New("Invalid Version ")
		return
	}

	if strings.TrimSpace(c.Password) == "" {
		err = errors.New("Invalid Version ")
		return
	}
	return
}
