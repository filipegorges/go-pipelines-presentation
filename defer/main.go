// START 1 OMIT
package main

import "fmt"

func defersAreSimilarToTryWithResources() bool {
	defer fmt.Println("I'll print before return!!") // HL

	usages := []string{
		"Closing a file",
		"Closing a connection",
		"Closing a channel (SPOILERS)",
		"General resource clean-up", // HL
	}

	for _, example := range usages {
		fmt.Println(example)
	}

	return true // HL
}
// END 1 OMIT
// START 2 OMIT
func main() {
	if defersAreSimilarToTryWithResources() {
		fmt.Println("Let's move on!") // HL
	}
}
// END 2 OMIT
