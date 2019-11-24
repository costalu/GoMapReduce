package main

import (
	"fmt"
	"sort"
)

func isPreposition(word string) bool {
	switch word {
	case
		"a",
		"an",
		"the",
		"and",
		"or",
		"but",
		"etc",
		"for",
		"in",
		"to",
		"is",
		"so",
		"are",
		"on",
		"of",
		"this",
		"that",
		"with",
		"it",
		"like":
		return true
	}
	return false
}

func printMostFrequentWords(m map[string]int) {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for i, kv := range ss {
		fmt.Printf("%s -> %d\n", kv.Key, kv.Value)
		if i > 100 {
			break
		}
	}
}
