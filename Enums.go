// package main identifies this file as an executable program
package main

// import "fmt" allows us to print text and errors
import "fmt"

// Define a new type 'ServerState' based on an integer (int).
// This makes our code easier to read than just using numbers 0, 1, 2...
type ServerState int

// Use 'const' to define our fixed states.
// 'iota' is a special Go keyword that automatically assigns numbers starting from 0.
const (
	StateIdle      ServerState = iota // StateIdle = 0
	StateConnected                    // StateConnected = 1
	StateError                        // StateError = 2
	StateRetrying                     // StateRetrying = 3
)

// Create a map to link the numeric states to human-readable words (strings).
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// This is a method for the ServerState type.
// It allows the 'fmt' package to print the name (like "idle") instead of the number (0).
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	// Start at 'StateIdle' and move to the next state using the transition function
	ns := transition(StateIdle)
	// Prints "connected" because transition(StateIdle) returns StateConnected
	fmt.Println(ns)

	// Take the current state (connected) and transition again
	ns2 := transition(ns)
	// Prints "idle" because transition(StateConnected) returns StateIdle
	fmt.Println(ns2)
}

// This function defines the "logic" of our server.
// It decides what the next state should be based on the current state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		// If we are idle, we try to connect
		return StateConnected
	case StateConnected, StateRetrying:
		// If we are already connected or retrying, we go back to idle
		return StateIdle
	case StateError:
		// If there is an error, stay in the error state
		return StateError
	default:
		// If the state is something we don't recognize, stop the program (panic)
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
