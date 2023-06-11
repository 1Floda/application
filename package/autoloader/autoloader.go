package autoloader

import (
	"github.com/joho/godotenv"
	"log"
)

func AutoLoad() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Ошибка загрузки ENV")
	}
}
