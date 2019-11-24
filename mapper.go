package main

import "strings"

// Map returns a map of word -> frequency for a given string
func Map(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := map[string]int{}
	for _, word := range wordList {
		lowerWord := strings.ToLower(word)

		if !isPreposition(lowerWord) {
			_, exists := counts[lowerWord]
			if exists {
				counts[lowerWord]++
			} else {
				counts[lowerWord] = 1
			}
		}
	}
	return counts
}
