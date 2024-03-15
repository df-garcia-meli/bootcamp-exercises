package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal"
	"os"
)

type ProductMap struct {
	Products map[int]internal.Product `json:"products"`
}

func NewProductMap() *ProductMap {
	// Open our jsonFile
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened json as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	var products []internal.Product
	json.Unmarshal(byteValue, &products)

	// Map products by id
	productsMap := make(map[int]internal.Product)
	for _, product := range products {
		productsMap[product.Id] = product
	}

	return &ProductMap{
		Products: productsMap,
	}
}

func (p *ProductMap) Save(product *internal.Product) (err error) {
	product.Id = len(p.Products) + 1
	p.Products[len(p.Products)+1] = *product
	return nil
}

func (p *ProductMap) Update(product *internal.Product) (err error) {
	p.Products[product.Id] = *product
	return nil
}

func (p *ProductMap) UpdatePartial(id int, fields map[string]any) (err error) {
	product, ok := p.Products[id]
	if !ok {
		return internal.ErrProductNotFound
	}

	for key, value := range fields {
		switch key {
		case "name":
			product.Name = value.(string)
		case "quantity":
			product.Quantity = value.(int)
		case "code_value":
			product.CodeValue = value.(string)
		case "is_published":
			product.IsPublished = value.(bool)
		case "expiration":
			product.Expiration = value.(string)
		case "price":
			product.Price = value.(float64)
		}
	}

	p.Products[id] = product
	return nil
}

func (p *ProductMap) Delete(id int) (err error) {
	_, ok := p.Products[id]
	if !ok {
		return internal.ErrProductNotFound
	}

	delete(p.Products, id)
	return nil
}
