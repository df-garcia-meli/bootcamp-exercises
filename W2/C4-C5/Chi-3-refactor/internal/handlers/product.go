package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"main/internal"
	"main/platform/tools"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Product struct {
	sv internal.ProductService
}

func NewProduct(sv internal.ProductService) *Product {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Product{
		sv: sv,
	}
}

func (p *Product) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get headers
		header_token := r.Header.Get("TOKEN")
		secret_token := os.Getenv("TOKEN")
		if header_token != secret_token {
			response.JSON(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Get request bytes
		requestBytes, _ := io.ReadAll(r.Body)

		// Unmarshal request bytes to map
		var requestMap map[string]any
		json.Unmarshal(requestBytes, &requestMap)

		// Validate request
		if err := tools.CheckFieldExistance(requestMap, "name", "quantity", "code_value", "expiration", "price"); err != nil {
			response.JSON(w, http.StatusBadRequest, err)
		}

		// Bytes to RequestProduct
		var requestProduct RequestProduct
		json.Unmarshal(requestBytes, &requestProduct)

		// // Check unique code_value
		// for _, product := range s.Products {
		// 	if product.CodeValue == requestProduct.CodeValue {
		// 		code := http.StatusBadRequest
		// 		w.WriteHeader(code)
		// 		w.Write([]byte("Code value must be unique"))
		// 		return
		// 	}
		// }

		// // Check date
		// _, err := time.Parse("01/02/2006", requestProduct.Expiration)
		// if err != nil {
		// 	code := http.StatusBadRequest
		// 	w.WriteHeader(code)
		// 	w.Write([]byte("Invalid date"))
		// 	return
		// }

		// Create product based on requestProduct
		newProduct := internal.Product{
			Name:        requestProduct.Name,
			Quantity:    requestProduct.Quantity,
			CodeValue:   requestProduct.CodeValue,
			IsPublished: requestProduct.IsPublished,
			Expiration:  requestProduct.Expiration,
			Price:       requestProduct.Price,
		}
		if err := p.sv.Save(&newProduct); err != nil {
			switch {
			case errors.Is(err, internal.ErrProducDuplicated):
				response.JSON(w, http.StatusConflict, err)
			default:
				response.JSON(w, http.StatusInternalServerError, err)
			}
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"product": newProduct,
		})
	}
}

func (p *Product) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get headers
		header_token := r.Header.Get("TOKEN")
		secret_token := os.Getenv("TOKEN")
		if header_token != secret_token {
			response.JSON(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Get productId from url
		productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

		// Get request bytes
		requestBytes, _ := io.ReadAll(r.Body)

		// Unmarshal request bytes to map
		var requestMap map[string]any
		json.Unmarshal(requestBytes, &requestMap)

		// Validate request
		if err := tools.CheckFieldExistance(requestMap, "name", "quantity", "code_value", "expiration", "price"); err != nil {
			response.JSON(w, http.StatusBadRequest, err)
		}

		// Bytes to RequestProduct
		var requestProduct RequestProduct
		json.Unmarshal(requestBytes, &requestProduct)

		// Create product based on requestProduct
		updatedProduct := internal.Product{
			Id:          productId,
			Name:        requestProduct.Name,
			Quantity:    requestProduct.Quantity,
			CodeValue:   requestProduct.CodeValue,
			IsPublished: requestProduct.IsPublished,
			Expiration:  requestProduct.Expiration,
			Price:       requestProduct.Price,
		}
		if err := p.sv.Update(&updatedProduct); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductInvalidField):
				response.JSON(w, http.StatusBadRequest, err)
			default:
				response.JSON(w, http.StatusInternalServerError, err)
			}
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"product": updatedProduct,
		})
	}
}

func (p *Product) UpdateProductPartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get headers
		header_token := r.Header.Get("TOKEN")
		secret_token := os.Getenv("TOKEN")
		if header_token != secret_token {
			response.JSON(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Get productId from url
		productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

		// Body into map
		var requestMap map[string]any
		json.NewDecoder(r.Body).Decode(&requestMap)

		if err := p.sv.UpdatePartial(productId, requestMap); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductInvalidField):
				response.JSON(w, http.StatusBadRequest, err)
			default:
				response.JSON(w, http.StatusInternalServerError, err)
			}
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product updated"))
	}
}

func (p *Product) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get headers
		header_token := r.Header.Get("TOKEN")
		secret_token := os.Getenv("TOKEN")
		if header_token != secret_token {
			response.JSON(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		// Get productId from url
		productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

		if err := p.sv.Delete(productId); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.JSON(w, http.StatusNotFound, err)
			default:
				response.JSON(w, http.StatusInternalServerError, err)
			}
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product deleted"))
	}
}
