//

package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var opr string
	fmt.Print("Enter opr: ")
	fmt.Scanln(&opr)
	fmt.Print("Enter num1: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter num2: ")
	fmt.Scanln(&num2)

	switch opr {
	case "+":
		var c float64 = num1 + num2
		fmt.Println("addition: ", c)
		break
	case "-":
		var d float64 = num1 - num2
		fmt.Println("subtraction: ", d)
		break
	case "*":
		var e float64 = num1 * num2
		fmt.Println("multiplication: ", e)
		break
	default:
		fmt.Println("Invalid operator")
	}
	// 	var c int = num1 + num2
	// 	var d int = num1 - num2
	// 	var e int = num1 * num2

	// 	fmt.Println(c)
	// 	fmt.Println(d)
	// 	fmt.Println(e)

}
