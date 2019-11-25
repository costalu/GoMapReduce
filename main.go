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

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

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

func distributed() {
	csvfile, err := os.Open("TweetsNBA.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	maps := make(chan map[string]int)

	length := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		length++

		go func(record []string) {
			maps <- Map(record[1])
		}(record)
	}

	wordFrequencies := DistributedReduce(maps, length)
	printMostFrequentWords(wordFrequencies)
}

func main() {
	start := time.Now()
	sequential()
	stop := time.Now()
	elapsedSequential := stop.Sub(start)

	start = time.Now()
	distributed()
	stop = time.Now()
	elapsedDistributed := stop.Sub(start)

	fmt.Println("Sequential Time: ", elapsedSequential)
	fmt.Println("Distributed Time: ", elapsedDistributed)
}
