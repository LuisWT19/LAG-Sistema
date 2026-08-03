package entities

import "time"

// Employee representa la información laboral de un empleado.
type Employee struct {
	BaseModel

	UserID uint `gorm:"not null;unique"`
	User   User `gorm:"foreignKey:UserID"`

	Position string    `gorm:"size:100;not null"`
	HireDate time.Time `gorm:"not null"`
}

// TableName especifica el nombre de la tabla.
func (Employee) TableName() string {
	return "employees"
}
