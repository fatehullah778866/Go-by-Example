// package main tells the Go compiler that this file should compile as an executable program
package main

// import "fmt" allows us to use functions like Println to show text on the screen
import "fmt"

func main() {

	// Create a slice (a list) of integers containing the numbers 2, 3, and 4
	nums := []int{2, 3, 4}

	// Create a variable named 'sum' and set it to 0 to keep track of the total
	sum := 0

	// Use 'range' to loop through the list.
	// The underscore (_) means we want to ignore the index (position) of the number.
	// 'num' represents the value of the current item in the list.
	for _, num := range nums {
		// Add the current number to the total sum
		sum += num
	}
	// Print the final result of the sum to the console
	fmt.Println("sum:", sum)

	// Loop through the list again, but this time we keep the index 'i'
	for i, num := range nums {
		// Check if the current number is equal to 3
		if num == 3 {
			// If it is 3, print its position (index) in the list
			fmt.Println("index:", i)
		}
	}

	// Create a map (a collection of keys and values).
	// Here, "a" is a key and "apple" is its value.
	kvs := map[string]string{"a": "apple", "b": "banana"}

	// Loop through the map. 'k' gets the key and 'v' gets the value.
	for k, v := range kvs {
		// Print the key and value in a formatted way (e.g., "a -> apple")
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Loop through the map again, but only take the keys 'k'
	for k := range kvs {
		// Print each key found in the map
		fmt.Println("key:", k)
	}

	// Loop through the string "go".
	// 'i' is the starting byte position and 'c' is the character code (Unicode point).
	for i, c := range "go" {
		// Print the position and the character code (71 for 'g', 111 for 'o')
		fmt.Println(i, c)
	}
}
