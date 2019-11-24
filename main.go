package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func sequential() {
	csvfile, err := os.Open("TweetsNBA.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	maps := []map[string]int{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		maps = append(maps, Map(record[1]))
	}

	wordFrequencies := Reduce(maps)
	printMostFrequentWords(wordFrequencies)
}

func concurrent() {
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
			maps <- Map(record[1])
		}(record)
	}

	wordFrequencies := FastReduce(maps, 5146)
	printMostFrequentWords(wordFrequencies)
}

func main() {
	start := time.Now()
	sequential()
	stop := time.Now()
	elapsedSequential := stop.Sub(start)

	start = time.Now()
	concurrent()
	stop = time.Now()
	elapsedConcurrent := stop.Sub(start)

	fmt.Println("Sequential Time: ", elapsedSequential)
	fmt.Println("Concurrent Time: ", elapsedConcurrent)
}
