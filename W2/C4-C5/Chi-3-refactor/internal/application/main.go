package main

import (
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Server
	router := chi.NewRouter()
	storeRepository := repository.NewProductMap()
	storeService := service.NewProductDefault(storeRepository)
	storeHandler := handlers.NewProduct(storeService)

	// Routes
	// router.Get("/ping", storeHandler.Ping())
	// router.Get("/products/search", storeHandler.GetProductsByPrice())
	// router.Get("/products/{productId}", storeHandler.GetProductById())
	// router.Get("/products", storeHandler.GetProducts())
	router.Post("/products", storeHandler.CreateProduct())
	router.Put("/products/{productId}", storeHandler.UpdateProduct())
	router.Patch("/products/{productId}", storeHandler.UpdateProductPartial())
	router.Delete("/products/{productId}", storeHandler.DeleteProduct())

	// Start server
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
