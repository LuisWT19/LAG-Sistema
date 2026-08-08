package handlers

import (
	"fmt"
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

// Create crea una nueva categoría.
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

// FindAll obtiene todas las categorías.
func (h *CategoryHandler) FindAll(c *gin.Context) {

	categories, err := h.service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

// FindByID obtiene una categoría por su ID.
func (h *CategoryHandler) FindByID(c *gin.Context) {

	idParam := c.Param("id")

	var id uint

	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de categoría inválido",
		})
		return
	}

	category, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoría no encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

// Update modifica una categoría existente.
func (h *CategoryHandler) Update(c *gin.Context) {

	idParam := c.Param("id")

	var id uint

	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de categoría inválido",
		})
		return
	}

	var request dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoría no encontrada",
		})
		return
	}

	category.Name = request.Name
	category.Description = request.Description
	category.Slug = request.Slug
	category.IsActive = request.IsActive

	if err := h.service.Update(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoría actualizada correctamente",
		"data":    category,
	})
}

// Delete elimina una categoria existente
func (h *CategoryHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")

	var id uint

	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de categoría inválido",
		})
		return
	}

	category, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Categoría no encontrada",
		})
		return
	}

	if err := h.service.Delete(category.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoría eliminada correctamente",
	})
}
