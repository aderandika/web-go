package main

import (
	"go-web/config"
	"go-web/controllers/CategoryController"
	"go-web/controllers/HomeController"
	"go-web/controllers/ProductController"
	"log"
	"net/http"
)

func main() {
	// Memanggil Database
	config.ConnectDB()

	// HomePage
	http.HandleFunc("/", HomeController.Welcome)

	// Categories
	http.HandleFunc("/categories", CategoryController.Index)
	http.HandleFunc("/categories/add", CategoryController.Add)
	http.HandleFunc("/categories/edit", CategoryController.Edit)
	http.HandleFunc("/categories/delete", CategoryController.Delete)

	// Product
	http.HandleFunc("/products", ProductController.Index)
	http.HandleFunc("/products/add", ProductController.Add)
	http.HandleFunc("/products/edit", ProductController.Edit)
	http.HandleFunc("/products/delete", ProductController.Delete)

	log.Println("Server sedang berjalan")
	http.ListenAndServe(":8080", nil)
}
