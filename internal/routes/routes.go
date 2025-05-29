package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goMarket/internal"
	"goMarket/internal/handlers"
	"goMarket/internal/middleware"
	"goMarket/internal/services"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateRouters(router *gin.Engine, db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	// Инициализация хранилища и сервисов продукта
	productService := services.NewProductService(db)
	productHandler := handlers.NewProductHandler(productService)

	// Инициализация хранилища и сервисов категории
	categoryService := services.NewCategoryService(db)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	shoppingCartService := services.NewShoppingCartService(db)
	shoppingCartHandler := handlers.NewShoppingCartHandler(shoppingCartService)

	cartItemService := services.NewCartItemService(db)
	cartItemHandler := handlers.NewCartItemHandler(cartItemService)

	orderService := services.NewOrderService(db)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderItemService := services.NewOrderItemService(db)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemService)

	authService := services.NewJWTAuthService(internal.JWTSecretKey, db)
	authHandler := handlers.NewAuthHandler(authService)

	taskService := services.NewTaskService(productService)
	taskHandler := handlers.NewTaskHandler(taskService)

	router.Use(middleware.TimeoutMiddleware(15 * time.Second))

	// Настройка маршрута для Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		usersGroup := api.Group("/users")
		{
			usersGroup.Use(middleware.JWTMiddleware())
			usersGroup.GET("/", userHandler.GetAllUsers)
			usersGroup.GET("/:id", userHandler.GetUserByID)
			usersGroup.PUT("/:id", userHandler.UpdateCurrentUser)
			usersGroup.DELETE("/:id", userHandler.DeleteUser)

			usersGroup.GET("/me", userHandler.GetCurrentUser)
			usersGroup.PATCH("/me", userHandler.UpdateCurrentUser)
		}

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/jwt/login", authHandler.Login)
			authGroup.POST("/jwt/logout", middleware.JWTMiddleware(), authHandler.Logout)
			authGroup.POST("/register", userHandler.Register)
		}

		shoppingCarts := api.Group("/shopping_carts")
		{
			shoppingCarts.Use(middleware.JWTMiddleware())
			shoppingCarts.GET("/", shoppingCartHandler.GetAllShoppingCarts)
			shoppingCarts.GET("/:id", shoppingCartHandler.GetShoppingCartByID)
			shoppingCarts.POST("/", shoppingCartHandler.CreateShoppingCart)
			shoppingCarts.PUT("/:id", shoppingCartHandler.UpdateShoppingCart)
			shoppingCarts.DELETE("/:id", shoppingCartHandler.DeleteShoppingCart)
		}

		cartItems := api.Group("/cart_items")
		{
			cartItems.Use(middleware.JWTMiddleware())
			cartItems.GET("/", cartItemHandler.GetAllCartItems)
			cartItems.GET("/:id", cartItemHandler.GetCartItemByID)
			cartItems.POST("/", cartItemHandler.CreateCartItem)
			cartItems.PUT("/:id", cartItemHandler.UpdateCartItem)
			cartItems.DELETE("/:id", cartItemHandler.DeleteCartItem)
		}

		orders := api.Group("/orders")
		{
			orders.Use(middleware.JWTMiddleware())
			orders.GET("/", orderHandler.GetAllOrders)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.POST("/", orderHandler.CreateOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.DELETE("/:id", orderHandler.DeleteOrder)
		}

		orderItems := api.Group("/order_items")
		{
			orderItems.Use(middleware.JWTMiddleware())
			orderItems.GET("/", orderItemHandler.GetAllOrderItems)
			orderItems.GET("/:id", orderItemHandler.GetOrderItemByID)
			orderItems.POST("/", orderItemHandler.CreateOrderItem)
			orderItems.PUT("/:id", orderItemHandler.UpdateOrderItem)
			orderItems.DELETE("/:id", orderItemHandler.DeleteOrderItem)
		}

		taskRoutes := api.Group("/tasks")
		{
			taskRoutes.Use(middleware.JWTMiddleware())
			taskRoutes.POST("/", taskHandler.CreateTask)
			taskRoutes.GET("/:id", taskHandler.GetTaskStatus)
			taskRoutes.DELETE("/:id", taskHandler.CancelTask)
		}

	}

	// Добавляем логирование ко всем маршрутам.
	return middleware.RequestLogger(r)
}
