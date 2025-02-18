package main

import (
	"fmt"

	"github.com/tkmagesh/Cisco-GoFoundation-Feb-2025/11-modularity-app/math"
	/* "github.com/tkmagesh/Cisco-GoFoundation-Feb-2025/11-modularity-app/math/utils" */

	"github.com/fatih/color"
	ut "github.com/tkmagesh/Cisco-GoFoundation-Feb-2025/11-modularity-app/math/utils"
)

func main() {
	color.Yellow("App Executed!")
	// fmt.Println("App Executed!")
	fmt.Println(math.Add(100, 200))
	fmt.Println(math.Subtract(100, 200))
	fmt.Println("Operation Count :", math.OpCount())
	// fmt.Println("is 21 even? :", utils.IsEven(21))
	fmt.Println("is 21 even? :", ut.IsEven(21))
}
