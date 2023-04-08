package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	CodeValue  string  `json:"code_value"`
	IsPublised bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

var productsList = []Product{}

const filename = "products.json"

func main() {
	err := loadProducts(filename, &productsList)
	if err != nil {
		log.Printf("Error in load file %s: %s", filename, err.Error())
		return
	}

	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {c.String(200, "pong")})
	products := r.Group("/products")
	{
		products.GET("", Get())
		products.GET(":id", GetById())
		products.GET("/priceGt", FilterByPriceGt())
	}
	r.Run(":8080")
}

func loadProducts(path string, list *[]Product) error  {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		return err
	}
	return nil
}

func Get()gin.HandlerFunc  {
	return func(c *gin.Context)  {
		c.JSON(200, productsList)
	}
}

func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"Error":"Cound'n not cast id"})
			return
		}
		for _, product := range productsList{
			if product.Id == idInt {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"Error": "Not found"})
	}
}

func FilterByPriceGt() gin.HandlerFunc {
	return func(c *gin.Context){
		price := c.Query("priceGt")
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			c.JSON(400, gin.H{"Error":"Cound'n not cast id"})
			return
		}
		filterProducts := []Product{}
		for _,product := range productsList{
			if product.Price > priceFloat {
				filterProducts = append(filterProducts, product)
			}
		}
		c.JSON(200, filterProducts)
	}
}