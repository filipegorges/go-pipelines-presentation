package main

import "fmt"

func main() {
	// START 1 OMIT
	lines := []string{
		"We provide some data through '<-' on the right side",
		"and consume the data through '<-' on the left side.",
		"They are 'blocking' by nature, meaning you can't pass in more data",
		"without the consumer taking it out first, and vice-versa!",
		"They are unbuffered by default (pass 1, read 1, 1 at a time)",
		"but can also be buffered (pass/read as much as possible until chosen limit!).",
		"'Do not communicate by sharing memory; instead, share memory by communicating.'",
	}

	channelsAreLikeQueues := make(chan string, len(lines)) // HL
	defer close(channelsAreLikeQueues)                     // HL

	for i := 0; i < len(lines); i++ {
		channelsAreLikeQueues <- lines[i] // HL
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(<-channelsAreLikeQueues) // HL
	}
	// END 1 OMIT
}
