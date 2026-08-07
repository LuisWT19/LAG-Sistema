package repository

import (
	"github.com/LuisWT19/LAG-Sistema/internal/domain/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entities.Category) error
	FindAll() ([]entities.Category, error)
	FindByID(id uint) (*entities.Category, error)
	Update(category *entities.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
func (r *categoryRepository) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll() ([]entities.Category, error) {
	var categories []entities.Category

	err := r.db.Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (*entities.Category, error) {

	var category entities.Category

	err := r.db.First(&category, id).Error

	return &category, err
}

func (r *categoryRepository) Update(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}
