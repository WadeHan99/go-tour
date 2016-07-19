package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// initalize the wordMap variable
	wordMap := make(map[string]int)

	// separate the strings into words
	words := strings.Fields(s)

	// iterate over the words to count each instance
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
