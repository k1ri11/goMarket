package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goMarket/internal/routes"
	"goMarket/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	deployType := os.Getenv("DEPLOY_TYPE")
	var configFile string
	var configPath string
	if deployType == "remote" {
		configFile = "remote_config.yaml"
		configPath = filepath.Join("..", "..", configFile)
	} else {
		configPath = "configs/local_config.yaml"
	}

	config, err := pkg.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Не удалось загрузить конфигурацию:", err)
		return
	}
	fmt.Println(config)

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
