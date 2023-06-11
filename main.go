package main

import (
	"application/package/autoloader"
	"application/package/db"
	"application/package/models"
	"application/package/router"
)

func main() {
	autoloader.AutoLoad()

	db, err := db.Connect()

	if err != nil {
		panic("Не могу подключиться к  MYSQL")
	}

	db.AutoMigrate(&models.User{})

	router.InitRouter(db)
	router.Serve("localhost:8888")
}
