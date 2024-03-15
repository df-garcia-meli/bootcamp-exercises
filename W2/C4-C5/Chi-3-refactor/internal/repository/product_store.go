package repository

import (
	"encoding/json"
	"fmt"
	"main/internal"
	"os"
)

type ProductStore struct {
	ProductsFilePath string
}

func NewProductStore(filePath string) *ProductStore {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	return &ProductStore{
		ProductsFilePath: filePath,
	}
}

func (p *ProductStore) Save(product *internal.Product) (err error) {
	// Open our jsonFile
	jsonFile, err := os.Open(p.ProductsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Read the contents of the jsonFile
	fileInfo, _ := jsonFile.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = jsonFile.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	// Load products from JSON file
	var products []internal.Product
	json.Unmarshal(buffer, &products)

	newProduct := internal.Product{
		Id:          len(products) + 1,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
	products = append(products, newProduct)

	// Save products to JSON file
	jsonFile, err = os.Create(p.ProductsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(buffer)

	return nil
}

func (p *ProductStore) Update(product *internal.Product) (err error) {
	return nil
}

func (p *ProductStore) UpdatePartial(id int, fields map[string]any) (err error) {
	return nil
}

func (p *ProductStore) Delete(id int) (err error) {
	return nil
}
