package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	pr := GetProductRepository()
	prod1, _ := NewProduct("123", "Product description", 1234)
	prod2, _ := NewProduct("456", "Newer better product", 4567)
	pr.save(prod1)
	pr.save(prod2)

	lp, _ := pr.retrieve(prod1.Id)
	fmt.Printf("%v\n", lp)

	r := gin.Default()
	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")
		p, err := pr.retrieve(ProductId(id))
		if (err != nil) {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, p)
	})

	r.GET("/products", func(c *gin.Context) {
		p, _ := pr.retrieveAll()
		c.JSON(http.StatusOK, p)
	})

	r.POST("/product", func(c *gin.Context) {
		p := convertToProduct(c)

		jsonProduct, _ := pr.save(Product(p))
		c.JSON(http.StatusOK, jsonProduct)
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("Product to be deleted:", id)
		pr.delete(ProductId(id))
		c.Status(http.StatusOK)
	})

	r.Run("localhost:9000")
}
