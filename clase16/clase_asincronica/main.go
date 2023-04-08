package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id" binding:"required" `
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published" `
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

var products []Product

func main() {
	r := gin.Default()

	r.POST("/products", func(ctx *gin.Context) {
		var product Product
		ctx.BindJSON(&product)
		ctx.JSON(http.StatusOK, gin.H{"user": fmt.Sprint(product)})
	})
	r.Run(":8080")
	fmt.Println(products)
}