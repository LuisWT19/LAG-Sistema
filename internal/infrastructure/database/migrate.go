package database

import (
	"github.com/LuisWT19/LAG-Sistema/internal/domain/entities"
	"gorm.io/gorm"
)

// AutoMigrate crea o actualiza automáticamente las tablas del sistema.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Permission{},
		&entities.Role{},
		&entities.User{},
		// &entities.Customer{},
		// &entities.Employee{},
	)
}
