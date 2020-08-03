package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select { // HL
		case message1 := <-channel1: // HL
			fmt.Println("received", message1)
		case message2 := <-channel2: // HL
			fmt.Println("received", message2)
		}
	}
}

// END OMIT
