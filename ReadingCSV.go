package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"sort"
	"strings"
	"io"
	"log"
	"os"
)

func mapper(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, exists := counts[word]
        if exists {
            counts[word] += 1
        } else {
            counts[word] = 1
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

func main() {
	// Open the file
	csvfile, err := os.Open("TweetsNBA.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	maps := make([]map[string]int, 1)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		maps = append(maps, mapper(record[1]))
	}

	result := reducer(maps)

	// for word, count := range result{
	// 	if count > 500{
	// 	fmt.Println(word, count)
	// 	}
	// }

	type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range result {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    for i, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		if i > 100{
			break
		}
    }
}