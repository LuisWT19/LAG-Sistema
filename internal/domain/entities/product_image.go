package entities

// ProductImage representa una imagen asociada a un producto.
type ProductImage struct {
	BaseModel

	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`

	ImageURL string `gorm:"size:255;not null"`

	IsPrimary bool `gorm:"default:false"`
}

// TableName especifica el nombre de la tabla.
func (ProductImage) TableName() string {
	return "product_images"
}
