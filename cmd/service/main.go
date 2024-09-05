package main

import (
	"TestTask"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	var config *TestTask.Config
	var err error
	if config, err = TestTask.LoadConfig(); err != nil {
		log.Fatal().Err(err)
	}

	//создаем хранилище?//
	var main_store *store.Store
	if main_store, err = store.CreateStore(config.Database); err != nil {
		return
	}

}
