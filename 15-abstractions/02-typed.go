package main

import (
	"fmt"
	"math"
)

/* V1.0 */
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument is neither a Circle nor a Rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }: // any object with an Area() method
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument is neither a Circle nor a Rectangle")
	}
}
*/

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

/* Ver2.0 */
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

/* Ver3.0 */
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// contract composition
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      // x should be AreaFinder
	PrintPerimeter(x) // x should be PerimeterFinder
}

/*
func PrintShapeStats(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c) //Ver3.0
	*/
	PrintShapeStats(c)

	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r) //Ver3.0
	*/
	PrintShapeStats(r)

	s := Square{Side: 15}
	/*
		PrintArea(s)
		PrintPerimeter(s) //Ver3.0
	*/
	PrintShapeStats(s)

}
