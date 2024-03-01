package main

import "fmt"

func main() {
	var p Person = Person{
		Name:     "Juan",
		LastName: "Perez",
		DNI:      12345678,
		date:     12345678,
	}
	p.PersonDetail()

	product := ProductFactory("Laptop", 1000, "small")
	fmt.Println(product.Price())
}

type Person struct {
	Name     string
	LastName string
	DNI      int
	date     int
}

func (p Person) PersonDetail() {
	fmt.Println("Name", p.Name, "LastName", p.LastName, "DNI", p.DNI, "Date", p.date)
}

const (
	small  = "small"
	medium = "medium"
	large  = "large"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	Name     string
	BaseCost float64
}

func (s SmallProduct) Price() float64 {
	return s.BaseCost
}

type MediumProduct struct {
	Name     string
	BaseCost float64
}

func (m MediumProduct) Price() float64 {
	return m.BaseCost * 1.03
}

type LargeProduct struct {
	Name     string
	BaseCost float64
}

func (l LargeProduct) Price() float64 {
	return 2500 + l.BaseCost*1.06
}

func ProductFactory(name string, baseCost float64, size string) (p Product) {
	switch size {
	case small:
		p = SmallProduct{
			Name:     name,
			BaseCost: baseCost,
		}
	case medium:
		p = MediumProduct{
			Name:     name,
			BaseCost: baseCost,
		}
	case large:
		p = LargeProduct{
			Name:     name,
			BaseCost: baseCost,
		}
	}
	return
}
