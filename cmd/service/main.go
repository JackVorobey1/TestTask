package main

import (
	"TestTask/internal"
	"TestTask/store/adapters/postgresql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// Вызываем функцию Load из config.go
	internal.Load()
	postgresql.Init()
	db := postgresql.GetDB()

	if db != nil {
		fmt.Println("Подключение к базе данных успешно установлено.")
	} else {
		fmt.Println("Не удалось установить соединение с базой данных.")
	}

}
