package main

// Reduce reduces a slice of maps to a single map
func Reduce(maps []map[string]int) map[string]int {
	result := map[string]int{}

	for _, m := range maps {
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

// DistributedReduce reduces a channel of maps to a single map
func DistributedReduce(maps chan map[string]int, length int) map[string]int {
	result := map[string]int{}

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
