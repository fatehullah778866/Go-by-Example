// package main identifies this file as an executable program
package main

import (
	"fmt"  // Used for printing text
	"time" // Used for controlling time, like making the program wait
)

// f is a simple function that prints a string and a number three times
func f(from string) {
	// range 3 creates a loop that runs for 0, 1, and 2
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// This is a "direct" or "synchronous" call.
	// The program waits for this function to finish completely before moving to the next line.
	f("direct")

	// The 'go' keyword starts a new Goroutine.
	// This function will now run "in the background" concurrently.
	// The main program does NOT wait for it; it moves to the next line immediately.
	go f("goroutine")

	// You can also start a goroutine with an anonymous function (a function without a name).
	go func(msg string) {
		fmt.Println(msg)
	}("going") // We call it immediately by passing "going" in parentheses

	// We make the main program "sleep" for 1 second.
	// This is important because if the 'main' function ends, the whole program stops,
	// even if other goroutines are still working!
	time.Sleep(time.Second)

	// After waiting, we print "done"
	fmt.Println("done")
}
