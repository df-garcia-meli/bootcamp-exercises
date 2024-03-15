package internal

import "errors"

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

// Errors for repository
var (
	ErrProductNotFound     = errors.New("product not found")
	ErrProducDuplicated    = errors.New("product duplicated")
	ErrProductInvalidField = errors.New("product invalid field")
	ErrProductSaveConflic  = errors.New("product save conflict")
)

type ProductRepository interface {
	Save(p *Product) (err error)
	Update(p *Product) (err error)
	UpdatePartial(id int, fields map[string]any) (err error)
	Delete(id int) (err error)
}

// Errors for service
var (
	ErrProductService = errors.New("product service error")
)

type ProductService interface {
	Save(p *Product) (err error)
	Update(p *Product) (err error)
	UpdatePartial(id int, fields map[string]any) (err error)
	Delete(id int) (err error)
}
