package main

import (
	"fmt"
	"time"
)
// START OMIT
func f(from string) { // HL
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i) // HL
	}
}

func main() {
	f("direct") // HL

	go f("goroutine") // HL

	go func(msg string) {
		fmt.Println(msg)
	}("going") // HL

	time.Sleep(time.Second)
	fmt.Println("done") // HL
}
// END OMIT