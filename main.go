package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	pr := GetProductRepository()
	p, _ := NewProduct("123", "Product description", 1234)
	pr.save(p)

	lp, _ := pr.load(p.Id)
	fmt.Printf("%v\n", lp)

	r := gin.Default()
  	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Query("id")
		p, _ := pr.load(ProductId(id))
    	c.JSON(http.StatusOK, p)
  })
  r.Run()
}