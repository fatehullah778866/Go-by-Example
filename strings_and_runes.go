// package main identifies this file as an executable program
package main

import (
	"fmt"          // Used for printing text
	"unicode/utf8" // Used for advanced string/character handling
)

func main() {

	// Define a constant string with Thai characters meaning "Hello"
	const s = "สวัสดี"

	// len(s) counts the raw bytes in the string, not the number of letters.
	// In Thai, each character takes 3 bytes, so the length will be 18.
	fmt.Println("Len:", len(s))

	// This loop goes through every single byte in the string one by one.
	for i := 0; i < len(s); i++ {
		// %x prints the byte's value in hexadecimal (base 16)
		fmt.Printf("%x ", s[i])
	}
	fmt.Println() // Prints a new line

	// utf8.RuneCountInString correctly counts the actual number of characters (runes).
	// For "สวัสดี", this will correctly return 6.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Using 'range' on a string automatically processes it character by character.
	// 'idx' is the byte position where the character starts.
	// 'runeValue' is the Unicode value of the character.
	for idx, runeValue := range s {
		// %#U prints the character's Unicode code and its literal shape (like 'ส')
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	// This loop shows a manual way to decode a string character by character.
	// i = current position, w = width (how many bytes the character uses)
	for i, w := 0, 0; i < len(s); i += w {
		// DecodeRuneInString looks at the string and returns the first character and its byte width.
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)

		// Update 'w' so the loop knows how many bytes to jump forward
		w = width

		// Call a helper function to check what the character is
		examineRune(runeValue)
	}
}

// examineRune checks a single character (rune) and prints a message if it matches
func examineRune(r rune) {

	// Check if the character is the English letter 't'
	if r == 't' {
		fmt.Println("found tee")
		// Check if the character is the Thai letter 'ส'
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
