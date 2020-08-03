package main

import "fmt"

// START OMIT
func main() {
	landline := make(chan int) // HL
	answeringMachine := func() {
		for phoneCall := range landline { // HL
			fmt.Println("Calls from telemarketing: ", phoneCall)
		}
	}
	go answeringMachine() // HL

	telemarketing := func() {
		for phoneCall := 1; phoneCall < 5; phoneCall++ {
			landline <- phoneCall // HL
		}
	}
	telemarketing() // HL
}
// END OMIT