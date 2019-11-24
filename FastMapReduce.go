package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func mapper(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
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

func reducer(maps chan map[string]int, length int) map[string]int {
	result := make(map[string]int)

	for i := 0; i < length; i++ {
		m := <-maps
		for word, count := range m {
			_, exists := result[word]
			if exists {
				result[word] += count
			} else {
				result[word] = count
			}
		}
	}
	return result
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
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		if i > 100 {
			break
		}
	}
}

func main() {
	start := time.Now()

	csvfile, err := os.Open("TweetsNBA.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	maps := make(chan map[string]int)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		go func(record []string) {
			maps <- mapper(record[1])
		}(record)
	}

	wordFrequencies := reducer(maps, 5146)
	printMostFrequentWords(wordFrequencies)

	stop := time.Now()
	elapsed := stop.Sub(start)
	fmt.Println(elapsed)
}
