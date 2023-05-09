package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func convertToProduct(c *gin.Context) Product {
	var prodInput Product
	if err := c.ShouldBindJSON(&prodInput); err != nil {
		fmt.Println("Error:", err)
	}

	p, _ := NewProduct(prodInput.Id, prodInput.Description, prodInput.Price)
	return p
}