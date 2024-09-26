package store

import (
	"TestTask/store/adapters/pg"
	"errors"
)

type Config struct {
	Postgres *pg.Config `yaml:"postgres"`
}

func (c *Config) Validdate() (err error) {
	if c == nil {
		return errors.New("Empty store config")
	}
	if err = c.Postgres.Validdate(); err != nil {
		return
	}
	return
}
