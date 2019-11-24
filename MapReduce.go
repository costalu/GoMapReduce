package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func mapper(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		if !isPreposition(word) {
			_, exists := counts[word]
			if exists {
				counts[word]++
			} else {
				counts[word] = 1
			}
		}
	}
	return counts
}

func reducer(maps []map[string]int) map[string]int {
	res := make(map[string]int)

	for _, m := range maps {
		for word, count := range m {
			_, exists := res[word]
			if exists {
				res[word] += count
			} else {
				res[word] = count
			}
		}
	}
	return res
}

func isPreposition(word string) bool {
	switch word {
	case
		"a",
		"an",
		"the",
		"and",
		"or",
		"but",
		"etc":
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
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		if i > 100 {
			break
		}
	}
}

func main() {
	csvfile, err := os.Open("TweetsNBA.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	maps := make([]map[string]int, 1)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		maps = append(maps, mapper(record[1]))
	}

	wordFrequencies := reducer(maps)
	printMostFrequentWords(wordFrequencies)
}
