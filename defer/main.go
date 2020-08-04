package main

import "fmt"

// START 1 OMIT
func defersWillBeTheLastThingToRunBeforeReturning() bool {
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

func main() {
	if defersWillBeTheLastThingToRunBeforeReturning() {
		fmt.Println("Let's move on!") // HL
	}
}

// END 1 OMIT
