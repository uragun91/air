package main

import (
	avalid "air-api/air-valid"
	"air-api/auth"
	"air-api/database"
	"air-api/middlewares"
	"air-api/products"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// product represents data about a record product.
type Merchant struct {
	Id    string  `json:"id"`
	Name  string  `json:"title"`
}

func main() {
	err := godotenv.Load();
	if err != nil {
    log.Fatalf("Error loading .env file");
  }

	database.ConnectDB();
	database.RunMigrations();

	avalid.RegisterValidators();

	router := gin.Default();

	merchantsRouterGroup := router.Group("/api/merchants").Use(middlewares.JwtAuthMiddleware());
	customersRouterGroup := router.Group("/api");

	customersRouterGroup.GET("/product", products.GetProducts);
	customersRouterGroup.GET("/product/:id", products.GetProductById);

	merchantsRouterGroup.POST("/product", products.PostProduct);

	customersRouterGroup.POST("/register", auth.RegisterUser)

	port := os.Getenv("PORT");

	if port == "" {
		port = "8080";
	}

	if err := router.Run(":" + port); err != nil {
    log.Panicf("error: %s", err);
	}
}

