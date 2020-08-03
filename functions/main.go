// START 1 OMIT
package main

import "fmt"

func countWords(words []string, fn func(word string)) int { // HL
	wordCount := 0
	for _, word := range words {
		wordCount++
		fn(word) // HL
	}
	return wordCount
}
// END 1 OMIT
// START 2 OMIT
func main() {
	words := []string{"Closures", "are", "awesome", "!"}

	lenAcc := 0 // HL
	wordLenAcc := func(word string) { // HL
		lenAcc += len(word) // HL
	} // HL

	wc := countWords(words, wordLenAcc)

	fmt.Println("Number of words: ", wc)
	fmt.Println("Average length of words: ", lenAcc/wc) // HL
}
// END 2 OMIT