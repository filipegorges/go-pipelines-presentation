package main

import (
	"log"
)
func fromArray(arr []string) <-chan string {
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
func logger(name string, in <-chan string) <-chan string { // HL
	out := make(chan string) // HL
	go func() {
		defer close(out) // HL
		for str := range in {
			log.Printf("Logger %s received: %s", name, str) // HL
			out <- str // HL
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
	out = fromArray(wordsOfWisdom)
	out = logger("input", out) // HL
	// END MAINLOGGER1 OMIT
	for str := range out {
		log.Printf("Sink received: %s", str)
	}
}
// END MAIN OMIT