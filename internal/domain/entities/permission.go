package entities

//Permission representa una accion aurotizada dentro del sistema
type Permission struct {
	BaseModel

	Name        string `gorm:"size:100;unique;not null"`
	Code        string `gorm:"size:100;unique;not null"`
	Description string `gorm:"size:255"`
}
