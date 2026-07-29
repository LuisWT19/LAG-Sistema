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

	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 LAG API iniciada correctamente")
}
