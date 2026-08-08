package routes

import (
	"github.com/LuisWT19/LAG-Sistema/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func NewRoutes(router *gin.Engine) *Routes {
	return &Routes{
		router: router,
	}
}

func (r *Routes) RegisterCategoryRoutes(categoryHandler *handlers.CategoryHandler) {

	api := r.router.Group("/api/v1")

	{
		api.POST("/categories", categoryHandler.Create)
		api.GET("/categories", categoryHandler.FindAll)
		api.GET("/categories/:id", categoryHandler.FindByID)
		api.PUT("/categories/:id", categoryHandler.Update)
		api.DELETE("/categories/:id", categoryHandler.Delete)
	}
}
