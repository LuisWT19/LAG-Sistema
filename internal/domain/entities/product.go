package entities

//Product representa un producto del catalago.
type Product struct {
	BaseModel

	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm: "foreignkey:CategoryID"`

	SubcategoryID uint        `gorm: "not null"`
	Subcategory   Subcategory `gorm: "foreignkey:SubcategoryID"`

	Name           string `gorm: "size:150;not null"`
	ScientificName string `gorm: "size:150"`
	Description    string `gorm: "type:text"`

	SKU string `gorm:"size:50;unique;not null"`

	Color string `gorm:"size:50"`

	Unit string `gorm:"size:20;not null"`

	ProductType string `gorm:"size:20;not null"`

	CostPrice float64 `gorm:"type:decimal(10,2);not null"`

	SalePrice float64 `gorm:"type:decimal(10,2);not null"`

	IsCustomizable bool `gorm:"default:false"`

	IsActive bool `gorm:"default:true"`
}

// TableName especifica el nombre de la tabla.
func (Product) TableName() string {
	return "products"
}
