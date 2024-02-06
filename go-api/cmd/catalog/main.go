package main

import (
	"database/sql"
	"ecommerceDevFullCycle/internal/database"
	"ecommerceDevFullCycle/internal/service"
	"ecommerceDevFullCycle/internal/webserver"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerceDevFullCycle")
	if err != nil {
		//para a execução
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	//pega todos os logs do servidor
	c.Use(middleware.Logger)
	//quando da problema não cai o servidor
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductsByCategory)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
