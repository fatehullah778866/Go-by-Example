// package main identifies this file as an executable program
package main

import (
	"fmt"    // Used for printing output
	"iter"   // Used for the new Iterator types
	"slices" // Used for helper functions like Collect
)

// List is a generic Linked List (same as the previous example)
type List[T any] struct {
	head, tail *element[T]
}

// element is a single node in the Linked List
type element[T any] struct {
	next *element[T]
	val  T
}

// Push adds a new value to the end of the Linked List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All is an iterator method. It returns an iter.Seq[T].
// This allows you to use 'for range' directly on our custom List.
func (lst *List[T]) All() iter.Seq[T] {
	// We return a function that takes a 'yield' function as an argument
	return func(yield func(T) bool) {
		// We walk through the linked list
		for e := lst.head; e != nil; e = e.next {
			// 'yield' sends the current value to the for-loop.
			// If yield returns false, the loop has stopped (e.g., using 'break'),
			// so we must stop our function too.
			if !yield(e.val) {
				return
			}
		}
	}
}

// genFib is an iterator that generates Fibonacci numbers forever
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1 // Start the sequence with 1, 1

		for {
			// Yield the current number in the sequence
			if !yield(a) {
				return // Stop if the caller breaks the loop
			}
			// Calculate the next number: a becomes b, and b becomes a+b
			a, b = b, a+b
		}
	}
}

func main() {
	// Create a list of integers and push some values
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Because we defined All(), we can use range to iterate over the list
	for e := range lst.All() {
		fmt.Println(e)
	}

	// slices.Collect is a helper that takes an iterator
	// and turns it into a standard slice []int
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Using the Fibonacci generator
	for n := range genFib() {
		// Since genFib is an infinite loop, we need a 'break' condition
		if n >= 10 {
			break // This causes the 'yield' function in genFib to return false
		}
		fmt.Println(n)
	}
}
