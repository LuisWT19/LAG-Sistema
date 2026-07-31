package entities

// User representa una cuenta de acceso al sistema.
type User struct {
	BaseModel

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`

	Email string `gorm:"size:150;unique;not null"`

	Password string `gorm:"size:255;not null"`

	Phone string `gorm:"size:20"`

	IsActive bool `gorm:"default:true"`

	RoleID uint
	Role   Role `gorm:"foreignKey:RoleID"`
}

// TableName especifica el nombre de la tabla.
func (User) TableName() string {
	return "users"
}
