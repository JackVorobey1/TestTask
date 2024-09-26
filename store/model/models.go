package model

import "gorm.io/gorm"

type ExampleModel struct {
	gorm.Model
	ID    uint   `yaml:"id"`
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
