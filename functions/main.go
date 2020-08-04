package main

import (
	"fmt"
	"strings"
)

// START 1 OMIT
func MapInPlace(pieces []string, fn func(string) string) {
	for i := 0; i < len(pieces); i++ {
		pieces[i] = fn(pieces[i]) // HL
	}
}

func main() {
	words := []string{"Closures", "are", "awesome", "!"}

	// Closure as a variable // HL
	transformer := func(s string) string { // HL
		return strings.ToUpper(s) // HL
	} // HL

	MapInPlace(words, transformer)

	fmt.Printf("'words' in upper case: %#v", words)
}

// END 1 OMIT
