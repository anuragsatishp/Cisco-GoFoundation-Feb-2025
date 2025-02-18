package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product
	/* Dummy */
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

func main() {
	// milk := PerishableProduct{Product{100, "Milk", 50}, "2 Days"}
	/*
		milk := PerishableProduct{
			Product: Product{
				Id:   100,
				Name: "Milk",
				Cost: 50,
			},
			Expiry: "2 Days",
		}
	*/

	milk := NewPerishableProduct(100, "Milk", 50, "2 Days")
	// fmt.Printf("%#v\n", milk)
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)

	// Use Format() and ApplyDiscount() for milk (PerishableProduct)
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) /* no return values */ {
	/* apply the given discount on the given product */
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
