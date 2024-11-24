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

	shoppingCartService := services.NewShoppingCartService(db)
	shoppingCartHandler := handlers.NewShoppingCartHandler(shoppingCartService)

	cartItemService := services.NewCartItemService(db)
	cartItemHandler := handlers.NewCartItemHandler(cartItemService)

	orderService := services.NewOrderService(db)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderItemService := services.NewOrderItemService(db)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)

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

		shoppingCarts := api.Group("/shopping_carts")
		{
			shoppingCarts.GET("/", shoppingCartHandler.GetAllShoppingCarts)
			shoppingCarts.GET("/:id", shoppingCartHandler.GetShoppingCartByID)
			shoppingCarts.POST("/", shoppingCartHandler.CreateShoppingCart)
			shoppingCarts.PUT("/:id", shoppingCartHandler.UpdateShoppingCart)
			shoppingCarts.DELETE("/:id", shoppingCartHandler.DeleteShoppingCart)
		}

		cartItems := api.Group("/cart_items")
		{
			cartItems.GET("/", cartItemHandler.GetAllCartItems)
			cartItems.GET("/:id", cartItemHandler.GetCartItemByID)
			cartItems.POST("/", cartItemHandler.CreateCartItem)
			cartItems.PUT("/:id", cartItemHandler.UpdateCartItem)
			cartItems.DELETE("/:id", cartItemHandler.DeleteCartItem)
		}

		orders := api.Group("/orders")
		{
			orders.GET("/", orderHandler.GetAllOrders)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.POST("/", orderHandler.CreateOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.DELETE("/:id", orderHandler.DeleteOrder)
		}

		orderItems := api.Group("/order_items")
		{
			orderItems.GET("/", orderItemHandler.GetAllOrderItems)
			orderItems.GET("/:id", orderItemHandler.GetOrderItemByID)
			orderItems.POST("/", orderItemHandler.CreateOrderItem)
			orderItems.PUT("/:id", orderItemHandler.UpdateOrderItem)
			orderItems.DELETE("/:id", orderItemHandler.DeleteOrderItem)
		}

	}

	// Добавляем логирование ко всем маршрутам.
	return middleware.RequestLogger(r)
}
