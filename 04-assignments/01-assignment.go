/*
Create an interactive calculator

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit

if user input == 5
	then exit

if user input == 1 to 4
	accept the 2 operands from the user
	perform the corresponding operation
	print the result
	display the menu again

if user input < 1 OR > 5
	print "invalid choice"
	display the menu again

*/

package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		fmt.Println("1.Add")
		fmt.Println("2.Subtract")
		fmt.Println("3.Multiply")
		fmt.Println("4.Divide")
		fmt.Println("5.Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
