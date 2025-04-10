package v1

import (
	"inventory-service/internal/api/http/v1/handlers"
	"inventory-service/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(productUC *usecase.ProductUsecase, categoryUC *usecase.CategoryUsecase) *gin.Engine {
	r := gin.Default()

	productHandler := handlers.NewProductHandler(productUC)
	categoryHandler := handlers.NewCategoryHandler(categoryUC)

	// Продукты
	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", productHandler.List)
		productRoutes.POST("", productHandler.Create)
		productRoutes.GET("/:id", productHandler.GetByID)
		productRoutes.PATCH("/:id", productHandler.Update)
		productRoutes.DELETE("/:id", productHandler.Delete)
	}

	// Категории
	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.GET("", categoryHandler.List)
		categoryRoutes.POST("", categoryHandler.Create)
		categoryRoutes.GET("/:id", categoryHandler.GetByID)
		categoryRoutes.PATCH("/:id", categoryHandler.Update)
		categoryRoutes.DELETE("/:id", categoryHandler.Delete)
	}

	return r
}
