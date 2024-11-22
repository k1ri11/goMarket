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

	// Инициализация хранилища и сервисов категории
	categoryService := services.NewCategoryService(db)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	customerService := services.NewCustomerService(db)
	customerHandler := handlers.NewCustomerHandler(customerService)

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

		// Группа маршрутов для категорий
		categories := api.Group("/categories")
		{
			categories.GET("/", categoryHandler.GetAllCategories)
			categories.GET("/:id", categoryHandler.GetCategoryByID)
			categories.POST("/", categoryHandler.CreateCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// Группа маршрутов для покупателей
		customers := api.Group("/customers")
		{
			customers.GET("/", customerHandler.GetAllCustomers)
			customers.GET("/:id", customerHandler.GetCustomerByID)
			customers.POST("/", customerHandler.CreateCustomer)
			customers.PUT("/:id", customerHandler.UpdateCustomer)
			customers.DELETE("/:id", customerHandler.DeleteCustomer)
		}

	}

	// Добавляем логирование ко всем маршрутам.
	return middleware.RequestLogger(r)
}
