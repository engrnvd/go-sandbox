package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	count := map[string]int{}
	words := strings.Fields(s)

	for _, word := range words {
		count[word]++
	}

	return count
}

func main() {
	fmt.Println(WordCount("Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.\n\nYou might find strings.Fields helpful."))
}
