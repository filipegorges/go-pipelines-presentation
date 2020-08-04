package main

import (
	"fmt"
)

func createProducerFromArray(arr []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, str := range arr {
			out <- str
		}
	}()

	return out
}

// START LOGGER OMIT
func logToStackDriver(name string, in <-chan string) <-chan string { // HL
	out := make(chan string) // HL
	go func() {
		defer close(out) // HL
		for str := range in {
			fmt.Printf("StackDriver %q received: %s\n", name, str) // HL
			out <- str                                             // HL
		}
	}()
	return out // HL
}

//END LOGGER OMIT

// START MAIN OMIT
func main() {
	var out <-chan string

	wordsOfWisdom := []string{
		"Do not communicate by sharing memory;",
		"instead, share memory by communicating.",
	}

	// START MAINLOGGER1 OMIT
	out = createProducerFromArray(wordsOfWisdom)
	out = logToStackDriver("input", out) // HL
	// END MAINLOGGER1 OMIT
	for str := range out {
		fmt.Println("Consumer received: ", str)
	}
}

// END MAIN OMIT
