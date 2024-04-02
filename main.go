package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculator Go 1.21")
	fmt.Println("===================")
	fmt.Println("Which operation do you want to perform?(add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "add":
		fmt.Println("Result: ", num1+num2)
	case "subtract":
		fmt.Println("Result: ", num1-num2)
	case "multiply":
		fmt.Println("Result: ", num1*num2)
	case "divide":
		fmt.Println("Result: ", num1/num2)
	default:
		fmt.Println("Invalid operation")
	}

}
