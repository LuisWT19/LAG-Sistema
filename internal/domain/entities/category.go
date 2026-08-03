package entities

// Category representa una categoría principal del catálogo.
type Category struct {
	BaseModel

	Name        string `gorm:"size:100;not null;unique"`
	Description string `gorm:"size:255"`
	Slug        string `gorm:"size:120;not null;unique"`
	IsActive    bool   `gorm:"default:true"`
}

// TableName especifica el nombre de la tabla.
func (Category) TableName() string {
	return "categories"
}
