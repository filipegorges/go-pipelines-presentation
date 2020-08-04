package main

import "fmt"

// START FROMARRAY OMIT
func createProducerFromArray(arr []string) <-chan string { // HL
	out := make(chan string) // HL
	go func() {
		defer close(out) // HL
		for _, str := range arr {
			out <- str // HL
		}
	}()
	return out // HL
}

// END FROMARRAY OMIT
func main() {
	// START MAINFROMARRAY OMIT
	var out <-chan string // HL
	wordsOfWisdom := []string{
		"Do not communicate by sharing memory;",
		"instead, share memory by communicating.",
	}

	out = createProducerFromArray(wordsOfWisdom) // HL
	for str := range out {                       // HL
		fmt.Println("Consumer received: ", str)
	}
	// END MAINFROMARRAY OMIT
}
