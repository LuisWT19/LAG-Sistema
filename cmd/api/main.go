package main

import (
	"log"

	"github.com/LuisWT19/LAG-Sistema/internal/infrastructure/database"
	"github.com/LuisWT19/LAG-Sistema/internal/shared/config"
)

func main() {

	config.Load()

	cfg := config.Get()

	log.Println("===================================")
	log.Println(cfg.AppName)
	log.Println("===================================")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println(" Migraciones ejecutadas correctamente.")
	log.Println(" LAG API iniciada correctamente")
}
