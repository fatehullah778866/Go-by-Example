// package main identifies this file as an executable program
package main

// import "fmt" allows us to print the message we receive
import "fmt"

func main() {

	// Create a new channel using 'make'.
	// This channel is designed to carry 'string' data.
	// 'chan' is the keyword for channel.
	messages := make(chan string)

	// Start a new Goroutine using an anonymous function.
	go func() {
		// Use the arrow operator (<-) to send a value INTO the channel.
		// The syntax 'channel <- value' means "send value to channel".
		messages <- "ping"
	}()

	// Use the arrow operator (<-) on the left side to receive a value FROM the channel.
	// This line "blocks" (waits) until a message arrives.
	// The syntax 'variable := <-channel' means "receive from channel".
	msg := <-messages

	// Print the message we received ("ping")
	fmt.Println(msg)
}
