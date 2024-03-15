package main

import (
	"main/store"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Server
	router := chi.NewRouter()
	storeHandler := store.NewStore()

	// Routes
	router.Get("/ping", storeHandler.Ping())
	router.Get("/products/search", storeHandler.GetProductsByPrice())
	router.Get("/products/{productId}", storeHandler.GetProductById())
	router.Get("/products", storeHandler.GetProducts())
	router.Post("/products", storeHandler.CreateProduct())

	// Start server
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
