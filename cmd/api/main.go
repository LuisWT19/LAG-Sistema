package main

import (
	"fmt"
	"log"

	"github.com/LuisWT19/LAG-Sistema/internal/shared/config"
)

func main() {
	config.Load()

	cfg := config.Get()

	log.Println("================================")
	log.Println(cfg.AppName)
	log.Println("Ambiente:", cfg.AppEnv)
	log.Println("Puerto:", cfg.AppPort)
	log.Println("================================")

	fmt.Println("Proyecto Inciado Correctamente :)")

}
