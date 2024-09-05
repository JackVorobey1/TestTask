package model

import "TestTask/store/adapters/postgresql"

// Config структура для конфигурации модели
type Config struct {
	PsgQl *postgresql.Config `yaml:"psg_ql"`
}

// метод проверки на ошибки//
func (c *Config) Validdate() (err error) {
	if err = c.PsgQl.Validdate(); err != nil {
		return
	}
	return
}
