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

	// Получаем текущую рабочую директорию
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}
	deployType := os.Getenv("DEPLOY_TYPE")
	var configFile string
	if deployType == "remote" {
		configFile = "remote_config.yaml"
	} else {
		configFile = "local_config.yaml"
	}
	configPath := filepath.Join(wd, "configs", configFile)
	config, err := pkg.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Ошибка при загрузке конфигурации:", err)
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
