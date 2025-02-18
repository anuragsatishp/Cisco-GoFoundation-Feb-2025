package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* no return values */ {
	/* apply the given discount on the given product */
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

func main() {

	product := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Println("Before discount")

	// fmt.Println(Format(product))
	fmt.Println(product.Format())

	fmt.Println("Applying 10% discount")
	// ApplyDiscount(&product, 10)
	product.ApplyDiscount(10)

	fmt.Println("After discount")
	// fmt.Println(Format(product))
	fmt.Println(product.Format())

}
