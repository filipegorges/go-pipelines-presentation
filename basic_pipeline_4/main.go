package main

import (
	"fmt"
	"strings"
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
func logToStackDriver(name string, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for str := range in {
			fmt.Printf("StackDriver %q received: %s\n", name, str)
			out <- str
		}
	}()

	return out
}
func upper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for str := range in {
			out <- strings.ToUpper(str)
		}
	}()

	return out
}

func main() {

	var out <-chan string

	wordsOfWisdom := []string{
		"Do not communicate by sharing memory;",
		"instead, share memory by communicating.",
	}

	// START MAINLOGGER2 OMIT
	out = createProducerFromArray(wordsOfWisdom)
	out = logToStackDriver("input", out)
	out = upper(out)
	out = logToStackDriver("uppered", out) // HL
	// END MAINLOGGER2 OMIT

	for str := range out {
		fmt.Println("Consumer received: ", str)
	}
}
