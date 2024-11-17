package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"goMarket/internal/handlers"
	"goMarket/internal/middleware"
	"goMarket/internal/services"
	"gorm.io/gorm"
	"net/http"
)

func CreateRouters(router *gin.Engine, db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	// Инициализация хранилища и сервисов продукта
	productService := services.NewProductService(db)
	productHandler := handlers.NewProductHandler(productService)

	api := router.Group("/v1")
	{
		//Группа маршрутов для продуктов
		products := api.Group("/products")
		{
			products.GET("/:id", productHandler.GetProductByID)
			products.GET("/", productHandler.GetProducts)
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	// Добавляем логирование ко всем маршрутам.
	return middleware.RequestLogger(r)
}
