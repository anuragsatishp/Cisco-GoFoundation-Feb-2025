package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

/*
Create a "Sort" api (function) to sort the products by any attribute
	IMPORTANT: DO NOT write your own logic for sorting. instead use sort.Sort() function
*/

type Products []Product

func (products Products) String() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p))
	}
	return sb.String()
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}

// sort.Interface interface implementation
func (products Products) Len() int {
	return len(products)
}

// default comparison by "Id"
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// Sort by name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// Sort by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// Sort by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// Sort by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func main() {

	// sample data
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Default List")
	fmt.Println(products)

	fmt.Println("Sort (default [id])")
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort - By Name")
	/* byName := ByName{Products: products}
	sort.Sort(byName) */
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("Sort - By Cost")
	/* byCost := ByCost{Products: products}
	sort.Sort(byCost) */
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort - By Units")
	/* byUnits := ByUnits{Products: products}
	sort.Sort(byUnits) */
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort - By Category")
	/* byCategory := ByCategory{Products: products}
	sort.Sort(byCategory) */
	products.Sort("Category")
	fmt.Println(products)

}
