package main

import "fmt"

func multiply(x int, y int) int {
	return x * y
}
func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func divide(x int, y int) int {
	return x / y
}

func main() {
	var choice string
	var var1, var2 int
	fmt.Print("Choose a number: \n")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. EXIT CALC")
	fmt.Scan(&choice)
	for choice != "5" {
		fmt.Printf("\n\nEnter First Number:")
		fmt.Scan(&var1)
		fmt.Println("Enter Second Number:")
		fmt.Scan(&var2)
		fmt.Printf("\n")
		switch choice {
		case "1":
			fmt.Printf("Answer: %d", add(var1, var2))
			fmt.Printf("\n\nPress 5 to quit\n")
			fmt.Scan(&choice)
			break
		case "2":
			fmt.Printf("Answer: %d", sub(var1, var2))
			fmt.Printf("\n\nPress 5 to quit\n")
			fmt.Scan(&choice)
			break
		case "3":
			fmt.Printf("Answer: %d", multiply(var1, var2))
			fmt.Printf("\n\nPress 5 to quit\n")
			fmt.Scan(&choice)
			break
		case "4":
			fmt.Printf("Answer: %d", divide(var1, var2))
			fmt.Printf("\n\nPress 5 to quit\n")
			fmt.Scan(&choice)
			break
		case "5":
			choice = "5"
		}

	}
	fmt.Println("Thanks for using gocalc")
}
