package main // Defines the package.

import ( // Starts the import block.
	"fmt" // Executes a statement.
	"time" // Executes a statement.
) // Ends the import block.

func main() { // Defines the main function.
	i := 2 // Initializes a variable.
	fmt.Println("Write ", i, "as ") // Writes output to the console.
	switch i { // Starts a switch statement.
	case 1: // Defines a switch case.
		fmt.Println("one") // Writes output to the console.
	case 2: // Defines a switch case.
		fmt.Println("two") // Writes output to the console.
	case 3: // Defines a switch case.
		fmt.Println("three") // Writes output to the console.
	} // Ends the current block.
	switch time.Now().Weekday() { // Starts a switch statement.
	case time.Saturday, time.Sunday: // Defines a switch case.
		fmt.Println("It's the weekend") // Writes output to the console.
	default: // Defines the default switch case.
		fmt.Println("It's a weekday") // Writes output to the console.
	} // Ends the current block.

	t := time.Now() // Initializes a variable.
	switch { // Starts a switch statement.
	case t.Hour() < 12: // Defines a switch case.
		fmt.Println("It's before noon") // Writes output to the console.
	default: // Defines the default switch case.
		fmt.Println("It's after noon") // Writes output to the console.
	} // Ends the current block.

	whatAmI := func(i interface{}) { // Initializes a variable.
		switch t := i.(type) { // Initializes a variable.
		case bool: // Defines a switch case.
			fmt.Println("I'm a bool") // Writes output to the console.
		case int: // Defines a switch case.
			fmt.Println("I'm an int") // Writes output to the console.
		default: // Defines the default switch case.
			fmt.Println("Don't know type %T\n", t) // Writes output to the console.
		} // Ends the current block.
	} // Ends the current block.
	whatAmI(true) // Executes a statement.
	whatAmI(1) // Executes a statement.
	whatAmI("hey") // Executes a statement.
} // Ends the current block.
