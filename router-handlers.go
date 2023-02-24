package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getProducts responds with the list of all products as JSON
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// postProduct adds an product from JSON received in the request body.
func postProduct(c *gin.Context) {
	var newProduct Product

	// Call BindJSON to bind the received JSON to
	// newproduct.
	if err := c.BindJSON(&newProduct); err != nil {
			return
	}

	// Add the new product to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of products, looking for
	// an product whose ID value matches the parameter.
	for _, product := range products {
			if product.Id == id {
					c.IndentedJSON(http.StatusOK, product)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}