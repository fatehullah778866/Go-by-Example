// package main identifies this file as an executable program
package main

// import "fmt" allows us to print text and data structures
import "fmt"

// Define a new type called 'person'.
// It groups a string (name) and an integer (age) together.
type person struct {
	name string
	age  int
}

// newPerson is a "constructor" function.
// It creates a new person and returns a pointer (*person) to it.
func newPerson(name string) *person {

	// Initialize a person with the given name
	p := person{name: name}
	// Manually set the age to 42
	p.age = 42
	// Return the memory address of this person
	return &p
}

func main() {

	// Create a person by providing values in order: name then age
	fmt.Println(person{"Bob", 20})

	// Create a person by explicitly naming the fields (recommended way)
	fmt.Println(person{name: "Alice", age: 30})

	// Create a person with only a name. The age will default to 0.
	fmt.Println(person{name: "Fred"})

	// Create a person and immediately get its memory address (pointer)
	fmt.Println(&person{name: "Ann", age: 40})

	// Use our constructor function to create "Jon" (age will be 42)
	fmt.Println(newPerson("Jon"))

	// Assign a person struct to a variable 's'
	s := person{name: "Sean", age: 50}
	// Access a specific field using a dot (.)
	fmt.Println(s.name)

	// Create a pointer 'sp' that points to the struct 's'
	sp := &s
	// Go is smart: you can use the dot even with pointers to access fields
	fmt.Println(sp.age)

	// Change a value through the pointer. This changes the original 's'.
	sp.age = 51
	fmt.Println(sp.age)

	// Create an "anonymous struct."
	// This is a struct used only once without giving the type a name.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	// Print the anonymous dog struct
	fmt.Println(dog)
}
