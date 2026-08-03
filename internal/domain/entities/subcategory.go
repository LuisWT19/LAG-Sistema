package entities

// Subcategory representa una subcategoría del catálogo.
type Subcategory struct {
	BaseModel

	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	Slug        string `gorm:"size:120;not null"`
	IsActive    bool   `gorm:"default:true"`
}

// TableName especifica el nombre de la tabla.
func (Subcategory) TableName() string {
	return "subcategories"
}
