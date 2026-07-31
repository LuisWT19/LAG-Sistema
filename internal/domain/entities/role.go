package entities

// Role representa un conjunto de permisos asignados a un usuario.
type Role struct {
	BaseModel

	Name        string `gorm:"size:100;unique;not null"`
	Description string `gorm:"size:255"`

	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
