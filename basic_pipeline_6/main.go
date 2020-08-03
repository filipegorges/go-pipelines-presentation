package main

import (
	"fmt"
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
func quote(in <-chan string) <-chan string { // HL
	out := make(chan string)
	go func() {
		defer close(out) // HL
		for str := range in {
			out <- fmt.Sprintf("%q", str) // HL
		}
	}()
	return out // HL
}
func main() {
	var out <-chan string

	wordsOfWisdom := []string{
		"Do not communicate by sharing memory;",
		"instead, share memory by communicating.",
	}

	// START MAINLOGGER3 OMIT
	out = fromArray(wordsOfWisdom)
	out = logger("input", out)
	out = upper(out)
	out = logger("uppered", out)
	out = quote(out)
	out = logger("quoted", out) // HL
	// END MAINLOGGER3 OMIT

	for str := range out {
		log.Printf("Sink received: %s", str)
	}
}
