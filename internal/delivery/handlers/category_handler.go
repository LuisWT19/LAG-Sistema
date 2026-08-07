package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LuisWT19/LAG-Sistema/internal/application"
	"github.com/LuisWT19/LAG-Sistema/internal/delivery/dto"
	"github.com/LuisWT19/LAG-Sistema/internal/domain/entities"
)

type CategoryHandler struct {
	service application.CategoryService
}

func NewCategoryHandler(service application.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {

	var request dto.CreateCategoryRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category := entities.Category{
		Name:        request.Name,
		Description: request.Description,
		Slug:        request.Slug,
		IsActive:    request.IsActive,
	}

	err = h.service.Create(&category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Categoría creada correctamente",
		"data":    category,
	})
}
