// package main tells the Go compiler that this file is an executable
package main

// import "fmt" is used for formatting strings and printing to the console
import "fmt"

// Define a simple struct called 'base' with one integer field
type base struct {
	num int
}

// Define a method for the 'base' struct.
// It returns a string that describes the 'num' value.
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Define a struct called 'container'.
// By typing 'base' without a name, we "embed" it.
// This means 'container' now has all the fields and methods of 'base'.
type container struct {
	base // This is the embedded struct
	str  string
}

func main() {

	// Create a new 'container' object.
	// We must initialize the inner 'base' struct explicitly.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Because of embedding, we can access 'num' directly on 'co'.
	// We don't have to type co.base.num (though we can).
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// This shows that we can still access the field through the full path if we want
	fmt.Println("also num:", co.base.num)

	// The 'container' also "borrows" the describe() method from 'base'.
	fmt.Println("describe:", co.describe())

	// Define an interface locally called 'describer'.
	// Anything with a describe() method satisfies this interface.
	type describer interface {
		describe() string
	}

	// Since 'container' has the describe() method (via embedding),
	// we can assign 'co' to the interface variable 'd'.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
