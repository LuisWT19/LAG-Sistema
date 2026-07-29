package database

import (
	"fmt"
	"log"

	"github.com/LuisWT19/LAG-Sistema/internal/shared/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database respresentada la conexión con la base de datos
type Database struct {
	DB *gorm.DB
}

// Connect establece la conexión con MySQL
func Connect() (*Database, error) {

	cfg := config.Get()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Base de datos conectada correctamente")

	return &Database{
		DB: db,
	}, nil
}
