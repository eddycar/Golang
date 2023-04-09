package main

import (
	"database/sql"
	"net/http"

	"repositoryapi/cmd/server/handler"
	"repositoryapi/internal/product"
	"repositoryapi/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:33060)/my_db")

	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	storage := store.SqlStore{db}
	repo := product.Repository{&storage}
	serv := product.Service{&repo}
	prodHandler := handler.ProductHandler{&serv}

	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) {ctx.String(http.StatusOK, "pong")})
	productGroup :=  r.Group("/products")
	{
		productGroup.GET(":id", prodHandler.GetById)
	}
	
	r.Run(":8080")
}
