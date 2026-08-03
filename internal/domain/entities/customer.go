package entities

// Customer representa la información comercial de un cliente.
type Customer struct {
	BaseModel

	UserID uint `gorm:"not null;unique"`
	User   User `gorm:"foreignKey:UserID"`

	LoyaltyPoints int    `gorm:"default:0"`
	Notes         string `gorm:"size:500"`
}

// TableName especifica el nombre de la tabla.
func (Customer) TableName() string {
	return "customers"
}
