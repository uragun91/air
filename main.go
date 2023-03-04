package main

import (
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

	router := gin.Default();

	privateRouterGroup := router.Group("/api/private").Use(middlewares.JwtAuthMiddleware());
	publicRouterGroup := router.Group("/api");

	publicRouterGroup.GET("/product", products.GetProducts);
	publicRouterGroup.GET("/product/:id", products.GetProductById);
	privateRouterGroup.POST("/product", products.PostProduct);

	port := os.Getenv("PORT");

	if port == "" {
		port = "8080";
	}

	if err := router.Run(":" + port); err != nil {
    log.Panicf("error: %s", err);
	}
}

