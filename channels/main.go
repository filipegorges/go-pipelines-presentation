// START 1 OMIT
package main

import "fmt"

func main() {
	lines := []string{
		"We provide some data through '<-' on the right side", // HL
		"and consume the data through '<-' on the left side.", // HL
		"They are 'blocking' by nature, meaning you can't pass in more data", // HL
		"without the consumer taking it out first, and vice-versa!",
		"They are unbuffered by default (pass 1, read 1, 1 at a time)",
		"but can also be buffered (pass/read as much as possible until chosen limit!).",
		"'Do not communicate by sharing memory; instead, share memory by communicating.'", // HL
	}
	// END 1 OMIT
	// START 2 OMIT
	channelsAreLikeQueues := make(chan string, len(lines)) // HL
	defer close(channelsAreLikeQueues) // HL

	for i := 0; i < len(lines); i++ {
		channelsAreLikeQueues <- lines[i] // HL
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(<-channelsAreLikeQueues) // HL
	}
}
// END 2 OMIT

