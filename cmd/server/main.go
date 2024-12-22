package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goMarket/docs"
	"goMarket/internal/routes"
	"goMarket/internal/utils"
	"goMarket/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// @title           goMarket API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @tokenUrl /v1/auth/jwt/login
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	docs.SwaggerInfo.Host = utils.GetDynamicHost()

	deployType := os.Getenv("DEPLOY_TYPE")
	var configFile string
	var configPath string
	if deployType == "remote" {
		configFile = "configs/remote_config.yaml"
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
