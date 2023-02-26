package main

import (
	"air-api/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// product represents data about a record product.
type Merchant struct {
	Id    string  `json:"id"`
	Name  string  `json:"title"`
}

type Product struct {
	Id	string `json:"id"`
	Name string `json:"name"`
	Price uint `json:"price"`
	ImgSrc string `json:"imageSrc"`
}

var products = []Product{
	{Id: "1", Name: "Blue Train", Price: 400000, ImgSrc: "https://picsum.photos/id/1/600/800"},
	{Id: "2", Name: "Blue Train", Price: 111000, ImgSrc: "https://picsum.photos/id/2/600/800"},
	{Id: "3", Name: "Blue Train", Price: 5565600, ImgSrc: "https://picsum.photos/id/3/600/800"},
	{Id: "4", Name: "Blue Train", Price: 300000, ImgSrc: "https://picsum.photos/id/4/600/800"},
}

func main() {
	router := gin.Default()

	protectedRouterGroup := router.Group("/api").Use(middlewares.JwtAuthMiddleware());

	protectedRouterGroup.GET("/product", getProducts);
	protectedRouterGroup.POST("/product", postProduct);
	protectedRouterGroup.GET("/product/:id", getProductById);

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
    log.Panicf("error: %s", err)
	}
}

