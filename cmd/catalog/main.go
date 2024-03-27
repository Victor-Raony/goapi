package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Victor-Raony/goapi.git/internal/entity/database"
	"github.com/Victor-Raony/goapi.git/internal/service"
	"github.com/Victor-Raony/goapi.git/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", WebCategoryHandler.GetCategory)
	c.Get("/category", WebCategoryHandler.GetCategories)
	c.Post("/category", WebCategoryHandler.CreateCategory)

	c.Get("/product/{id}", WebProductHandler.GetProduct)
	c.Get("/product", WebProductHandler.GetProducts)
	c.Get("/product/category/{id}", WebProductHandler.GetProductByCategoryID)
	c.Post("/product", WebProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)

}
