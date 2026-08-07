package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/LuisWT19/LAG-Sistema/internal/application"
	"github.com/LuisWT19/LAG-Sistema/internal/delivery/handlers"
	"github.com/LuisWT19/LAG-Sistema/internal/delivery/routes"
	"github.com/LuisWT19/LAG-Sistema/internal/infrastructure/database"
	"github.com/LuisWT19/LAG-Sistema/internal/infrastructure/repository"
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

	// Crear el router de Gin
	router := gin.Default()

	// Repository
	categoryRepository := repository.NewCategoryRepository(db)

	// Service
	categoryService := application.NewCategoryService(categoryRepository)

	// Handler
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Routes
	appRoutes := routes.NewRoutes(router)
	appRoutes.RegisterCategoryRoutes(categoryHandler)

	log.Println("Migraciones ejecutadas correctamente.")
	log.Println("LAG API iniciada correctamente")

	err = router.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal(err)
	}
}
