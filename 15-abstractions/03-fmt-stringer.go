package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

// fmt.Stringer interface implementation
func (p Product) String() string {
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

// Struct composition
type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func main() {

	product := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Println("Before discount")
	fmt.Println(product)

	fmt.Println("Applying 10% discount")
	product.ApplyDiscount(10)

	fmt.Println("After discount")
	fmt.Println(product)

	// Using PerishableProduct
	fmt.Println("Using Perishable Product")
	milk := NewPerishableProduct(100, "Milk", 50, "2 Days")
	// Use Format() and ApplyDiscount() for milk (PerishableProduct)
	fmt.Println("before discount")
	fmt.Println(milk)
	milk.ApplyDiscount(10)
	fmt.Println("after discount")
	fmt.Println(milk)

}
