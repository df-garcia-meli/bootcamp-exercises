package store

import "net/http"

type StoreInterface interface {
	Ping() http.HandlerFunc
	GetProducts() http.HandlerFunc
	GetProductById() http.HandlerFunc
	GetProductsByPrice() http.HandlerFunc
	CreateProduct() http.HandlerFunc
}
