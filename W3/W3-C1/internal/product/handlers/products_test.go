package handlers_test

import (
	"app/internal/auth"
	"app/internal/product/handlers"
	"app/internal/product/repository"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	t.Run("should return a list of products", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.Get()

		// When
		req := httptest.NewRequest("GET", "/products", nil)
		req.Header.Set("Token", "SECRET_TOKEN")
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"message": "products", "data":[
			{
				"id":1,"name":"Product 1","quantity":10,"code_value":"code1","is_published":true,"expiration":"2021-12-31","price":100
			},
			{
				"id":2,"name":"Product 2","quantity":20,"code_value":"code2","is_published":true,"expiration":"2021-12-31","price":200
			}]}`, res.Body.String())
	})

	t.Run("should return a product", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.GetByID()

		// When
		req := httptest.NewRequest("GET", "/products/1", nil)

		// Add path param
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		// Add Auth header
		req.Header.Set("Token", "SECRET_TOKEN")

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"message": "product", "data":{
				"id":1,"name":"Product 1","quantity":10,"code_value":"code1","is_published":true,"expiration":"2021-12-31","price":100
			}}`, res.Body.String())
	})

	t.Run("should fail returning a product - invalid id string", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.GetByID()

		// When
		req := httptest.NewRequest("GET", "/products/1", nil)

		// Add path param
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "INVALID")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		// Add Auth header
		req.Header.Set("Token", "SECRET_TOKEN")

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"status": "Bad Request", "message":"Invalid id"}`, res.Body.String())
	})

	t.Run("should fail returning a product - id not found", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.GetByID()

		// When
		req := httptest.NewRequest("GET", "/products/1", nil)

		// Add path param
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "17")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		// Add Auth header
		req.Header.Set("Token", "SECRET_TOKEN")

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"status": "Not Found", "message":"Product not found"}`, res.Body.String())
	})

	t.Run("should fail returning a product - without token", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.GetByID()

		// When
		req := httptest.NewRequest("GET", "/products/1", nil)

		// Add path param
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		// Do Not Add Auth header
		//req.Header.Set("Token", "SECRET_TOKEN")

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusUnauthorized, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"status": "Unauthorized", "message":"Unauthorized"}`, res.Body.String())
	})
}

func TestPostProducts(t *testing.T) {
	t.Run("should create a new product", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.Create()

		// When
		req := httptest.NewRequest("POST", "/products", strings.NewReader(`{"name":"Product 3","quantity":30,"code_value":"code3","is_published":true,"expiration":"2021-12-31","price":300}`))
		req.Header.Set("Token", "SECRET_TOKEN")
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		require.JSONEq(t, `{"message": "product created", "data":{
			"id": 3, "name":"Product 3","quantity":30,"code_value":"code3","is_published":true,"expiration":"2021-12-31","price":300
		}}`, res.Body.String())
	})
}

func TestDeleteProducts(t *testing.T) {
	t.Run("should delete a product", func(t *testing.T) {
		// Given
		db := map[int]repository.ProductAttributesMap{
			1: {
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "code1",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
			2: {
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "code2",
				IsPublished: true,
				Expiration:  time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
				Price:       200.0,
			},
		}
		rp := repository.NewRepositoryProductMap(db, 2, time.DateOnly)
		hd := handlers.NewHandlerProducts(rp, auth.NewAuthTokenBasic("SECRET_TOKEN"))
		hdFunc := hd.Delete()

		// When
		req := httptest.NewRequest("DELETE", "/products", nil)

		// Add path param
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		// Add Auth header
		req.Header.Set("Token", "SECRET_TOKEN")
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Then
		require.Equal(t, http.StatusNoContent, res.Code)
		require.Equal(t, "", res.Header().Get("Content-Type"))
		require.Equal(t, "", res.Body.String())
	})

}
