package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/* product = id, name, cost */
	/*
		var product Product
		product.Id = 100
		product.Name = "Pen"
		product.Cost = 10
	*/

	/*
		var product Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	// var product Product = Product{100, "Pen", 10}
	/*
		var product Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	/*
		var product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	product := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	// fmt.Printf("%+v\n", product)
	fmt.Println(Format(product))

	p2 := product // a copy is created as struct instances are "values"
	p2.Cost = 200
	fmt.Printf("product.Cost = %v, p2.Cost = %v\n", product.Cost, p2.Cost)

	// struct instances are compared by value
	x := Product{99, "Stappler", 50}
	y := Product{99, "Stappler", 50}
	fmt.Println(x == y)

	// handling references using pointers
	var productPtr *Product
	productPtr = &product
	// Accessing members of a struct pointer using "." notation
	fmt.Println(productPtr.Id)

	fmt.Println("Cost before discount :", product.Cost)
	fmt.Println("Applying 10% discount")
	ApplyDiscount( /* ?....? */ )
	fmt.Println("Cost after discount :", product.Cost)
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount( /* ? */ ) /* no return values */ {
	/* apply the given discount on the given product */
}
