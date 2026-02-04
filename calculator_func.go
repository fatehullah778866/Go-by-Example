package main

import "fmt"

func calculate(a int, b int, op string) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		return a / b
	} else {
		fmt.Println("Invalid Operator")
		return 0
	}
}

func main() {
	var num1 int
	fmt.Print("Enter num1: ")
	fmt.Scan(&num1)
	var num2 int
	fmt.Print("Enter num2: ")
	fmt.Scan(&num2)
	var operator string
	fmt.Print("Enter operator: ")
	fmt.Scan(&operator)
	result := calculate(num1, num2, operator)
	fmt.Println("Result: ", result)
}
