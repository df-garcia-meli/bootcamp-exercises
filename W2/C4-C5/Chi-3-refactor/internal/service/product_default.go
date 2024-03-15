package service

import "main/internal"

type ProductDefault struct {
	// Repository
	rp internal.ProductRepository
}

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {
	err = p.rp.Save(product)

	// Si se quiere agregar un error personalizado
	//if err != nil {
	//	err = internal.ErrProductService
	//}

	return
}

func (p *ProductDefault) Update(product *internal.Product) (err error) {
	return p.rp.Update(product)
}

func (p *ProductDefault) UpdatePartial(id int, fields map[string]any) (err error) {
	return p.rp.UpdatePartial(id, fields)
}

func (p *ProductDefault) Delete(id int) (err error) {
	return p.rp.Delete(id)
}
