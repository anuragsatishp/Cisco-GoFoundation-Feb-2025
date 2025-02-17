/*
Create an interactive calculator

When the program is executed:

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit

Accept the user choice
if user choice == 5
	then exit the application

if user choice == 1 to 4
	accept the 2 operands from the user
	perform the corresponding operation
	print the result
	display the menu again (IMPORTANT)

if user choice < 1 OR > 5
	print "invalid choice"
	display the menu again (IMPORTANT)

*/

package main

import "fmt"

func main() {
	var userChoice, x, y, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			fmt.Println("Done!")
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&x, &y)
		switch userChoice {
		case 1:
			result = x + y
		case 2:
			result = x - y
		case 3:
			result = x * y
		case 4:
			result = x / y
		}
		fmt.Println("Result :", result)
	}
}
