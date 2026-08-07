package application

import (
	"github.com/LuisWT19/LAG-Sistema/internal/domain/entities"
	"github.com/LuisWT19/LAG-Sistema/internal/infrastructure/repository"
)

type CategoryService interface {
	Create(category *entities.Category) error
	FindAll() ([]entities.Category, error)
	FindByID(id uint) (*entities.Category, error)
	Update(category *entities.Category) error
	Delete(id uint) error
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
	}
}

func (s *categoryService) Create(category *entities.Category) error {
	return s.repository.Create(category)
}

func (s *categoryService) FindAll() ([]entities.Category, error) {
	return s.repository.FindAll()
}

func (s *categoryService) FindByID(id uint) (*entities.Category, error) {
	return s.repository.FindByID(id)
}

func (s *categoryService) Update(category *entities.Category) error {
	return s.repository.Update(category)
}

func (s *categoryService) Delete(id uint) error {
	return s.repository.Delete(id)
}
