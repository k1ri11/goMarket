package main

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/routes"
	"goMarket/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	// Загрузка конфигурации
	config, err := pkg.LoadConfig("configs/config.yaml")
	if err != nil {
		println("Не удалось загрузить конфигурацию")
		return
	}

	databaseURL := pkg.GetDBUrl(config)
	println(databaseURL)

	// Инициализация базы данных
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Установка маршрутов
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Сервер работает на порту " + config.Server.Port,
		})
	})

	routes.CreateRouters(router, db)

	log.Println("Starting server on port", config.Server.Port)
	if err := http.ListenAndServe(":"+config.Server.Port, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
