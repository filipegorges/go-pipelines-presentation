package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepMs(delay int) {
	sleepMs := time.Duration(delay)
	time.Sleep(sleepMs * time.Millisecond)
}

func generateIntEvery(delayMs int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
			ch <- 200 - rand.Intn(100)
		}
	}()
	return ch
}

func generateFloatEvery(delayMs int) <-chan float32 {
	ch := make(chan float32)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
			ch <- rand.Float32()
		}
	}()
	return ch
}

// START OMIT
func main() {
	floatChannel := generateFloatEvery(200)
	intChannel := generateIntEvery(100)

	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case f := <-floatChannel: // HL
			fmt.Printf("float received: %.3f\n", f)
		case n := <-intChannel: // HL
			fmt.Printf("integer received: %d\n", n)
		case <-deadline: // HL
			fmt.Println("deadline reached [exiting]")
			return
		}
	}
}

// END OMIT
