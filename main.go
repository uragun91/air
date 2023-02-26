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

	protectedRouterGroup := router.Group("/api").Use(middlewares.JwtAuthMiddleware());

	protectedRouterGroup.GET("/product", products.GetProducts);
	protectedRouterGroup.POST("/product", products.PostProduct);
	protectedRouterGroup.GET("/product/:id", products.GetProductById);

	port := os.Getenv("PORT");

	if port == "" {
		port = "8080";
	}

	if err := router.Run(":" + port); err != nil {
    log.Panicf("error: %s", err);
	}
}

