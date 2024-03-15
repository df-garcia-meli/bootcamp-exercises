package store

import (
	"encoding/json"
	"fmt"
	"io"
	"main/product"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Store struct {
	Products map[int]product.Product `json:"products"`
}

func NewStore() *Store {
	// Open our jsonFile
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened json as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	var products []product.Product
	json.Unmarshal(byteValue, &products)

	// Map products by id
	productsMap := make(map[int]product.Product)
	for _, product := range products {
		productsMap[product.Id] = product
	}

	return &Store{
		Products: productsMap,
	}
}

func (s *Store) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func (s *Store) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s.Products)
	}
}

func (s *Store) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

		result, ok := s.Products[productId]
		if !ok {
			code := http.StatusNotFound
			w.WriteHeader(code)
			return
		}

		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func (s *Store) GetProductsByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt := r.URL.Query().Get("priceGt")

		parsedPrice, _ := strconv.ParseFloat(priceGt, 64)

		var result []product.Product
		for _, product := range s.Products {
			if product.Price >= parsedPrice {
				result = append(result, product)
			}
		}

		if len(result) == 0 {
			code := http.StatusNotFound
			w.WriteHeader(code)
			return
		}

		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (s *Store) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get request bytes
		requestBytes, _ := io.ReadAll(r.Body)

		// Unmarshal request bytes to map
		var requestMap map[string]any
		json.Unmarshal(requestBytes, &requestMap)

		// Validate request
		if _, ok := requestMap["name"]; !ok {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Name required"))
			return
		}

		if _, ok := requestMap["quantity"]; !ok {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Quantity required"))
			return
		}

		if _, ok := requestMap["code_value"]; !ok {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Name required"))
			return
		}

		if _, ok := requestMap["expiration"]; !ok {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Expiration required"))
			return
		}

		if _, ok := requestMap["price"]; !ok {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Price required"))
			return
		}

		// Bytes to RequestProduct
		var requestProduct RequestProduct
		json.Unmarshal(requestBytes, &requestProduct)

		// Check unique code_value
		for _, product := range s.Products {
			if product.CodeValue == requestProduct.CodeValue {
				code := http.StatusBadRequest
				w.WriteHeader(code)
				w.Write([]byte("Code value must be unique"))
				return
			}
		}

		// Check date
		_, err := time.Parse("01/02/2006", requestProduct.Expiration)
		if err != nil {
			code := http.StatusBadRequest
			w.WriteHeader(code)
			w.Write([]byte("Invalid date"))
			return
		}

		// Create product based on requestProduct
		newProduct := product.Product{
			Id:          len(s.Products) + 1,
			Name:        requestProduct.Name,
			Quantity:    requestProduct.Quantity,
			CodeValue:   requestProduct.CodeValue,
			IsPublished: requestProduct.IsPublished,
			Expiration:  requestProduct.Expiration,
			Price:       requestProduct.Price,
		}
		s.Products[newProduct.Id] = newProduct

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"product": newProduct,
		})
	}
}
