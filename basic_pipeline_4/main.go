package main

import (
	"log"
	"strings"
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
func logger(name string, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for str := range in {
			log.Printf("Logger %s received: %s", name, str)
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
	out = fromArray(wordsOfWisdom)
	out = logger("input", out)
	out = upper(out)
	out = logger("uppered", out) // HL
	// END MAINLOGGER2 OMIT

	for str := range out {
		log.Printf("Sink received: %s", str)
	}
}
