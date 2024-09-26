package main

import (
	"TestTask"
	"TestTask/store"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

var globalConfig *TestTask.Config

// Загрузка конфигурации, результат сохранятется в globalConfig
func init() {
	var err error
	globalConfig, err = TestTask.LoadConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func main() {
	var err error
	store.Default, err = store.InitStore(globalConfig.Store)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	//Уровень логирования контролирует, какие типы сообщений
	//будут выводиться в логи (DEBUG,ERROR и тд)
	log.Logger = log.With().Caller().Logger().Level(zerolog.DebugLevel).Output(zerolog.ConsoleWriter{Out: os.Stderr})

	/*router, err := NewRouter(globalConfig.Network)
	if err != nil {
		log.Fatal(err)
	}*/

	// Основная программа не завершится, пока не получит сигнал завершения
	sig := make(chan os.Signal, 1)
	log.Info().Msg("Ready")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info().Msg("Программа завершена")

}
