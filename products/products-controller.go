package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id			string 	`json:"id"`
	Name 		string 	`json:"name"`
	Price 	uint 		`json:"price"`
	ImgSrc 	string 	`json:"imageSrc"`
}

var products = []Product{
	{Id: "1", Name: "Blue Train", Price: 400000, ImgSrc: "https://picsum.photos/id/1/600/800"},
	{Id: "2", Name: "Green Train", Price: 111000, ImgSrc: "https://picsum.photos/id/2/600/800"},
	{Id: "3", Name: "Blue Train", Price: 5565600, ImgSrc: "https://picsum.photos/id/3/600/800"},
	{Id: "4", Name: "Blue Train", Price: 300000, ImgSrc: "https://picsum.photos/id/4/600/800"},
}


// getProducts responds with the list of all products as JSON
func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// postProduct adds an product from JSON received in the request body.
func PostProduct(c *gin.Context) {
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

func GetProductById(c *gin.Context) {
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