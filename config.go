package TestTask

import (
	"TestTask/store/adapters/postgresql"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func LoadConfig() (config *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile("config.yaml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return
	}
	if err = config.Validdate(); err != nil {
		return
	}
	config.Print()
	return
}

type Config struct {
	Version  string             `yaml:"version"`
	LogLevel string             `yaml:"log_level"`
	Database *postgresql.Config `yaml:"database"`
}

// метод Print, он выводит форматированные данные структуры
func (c *Config) Print() { fmt.Printf("%v\n", pretty.Formatter(c)) }

/*
	метод Validdate для структуры Config, он выводит ошибку если какое то из полей

содержит ошибку
*/
func (с *Config) Validdate() (err error) {
	if strings.TrimSpace(с.Version) == "" {
		err = errors.New("Invalid Version")
		return
	}
	if strings.TrimSpace(с.LogLevel) == "" {
		err = errors.New("Invalid LogLevel")
		return
	}

	if err = c.Database.Validdate(); err != nil {
		return
	}
	return
}
