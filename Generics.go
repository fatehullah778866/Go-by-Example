// package main identifies this file as an executable program
package main

// import "fmt" is used to print information to the console
import "fmt"

// SlicesIndex is a generic function.
// [S ~[]E, E comparable] means:
// 1. E is a type that can be compared using == (like string or int).
// 2. S is a slice containing elements of type E.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	// Loop through the slice
	for i := range s {
		// If the value matches the element, return the index
		if v == s[i] {
			return i
		}
	}
	// Return -1 if the value is not found in the slice
	return -1
}

// List is a generic "Linked List" struct.
// [T any] means it can hold a value of 'any' type (int, string, etc.).
type List[T any] struct {
	head, tail *element[T] // Points to the start and end of the list
}

// element represents a single "node" or "link" in our list.
type element[T any] struct {
	next *element[T] // Pointer to the next link
	val  T           // The actual data being stored
}

// Push is a method to add a new value to the end of the list.
func (lst *List[T]) Push(v T) {
	// If the list is empty (tail is nil)
	if lst.tail == nil {
		// Create the first element and set it as both head and tail
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		// Add a new element after the current tail
		lst.tail.next = &element[T]{val: v}
		// Move the tail marker to the new last element
		lst.tail = lst.tail.next
	}
}

// AllElements collects all values in the list and returns them as a slice.
func (lst *List[T]) AllElements() []T {
	var elems []T
	// Start at the head and follow the 'next' pointers until the end (nil)
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	// Create a slice of strings
	var s = []string{"foo", "bar", "zoo"}

	// Use the generic function. Go "infers" (guesses) that E is string.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// You can also call it manually by specifying the types: [[]string, string]
	_ = SlicesIndex[[]string, string](s, "zoo")

	// Create a new List specifically for integers (List[int])
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Print all numbers stored in the list
	fmt.Println("list:", lst.AllElements())
}
