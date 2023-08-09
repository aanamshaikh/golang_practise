package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	var result []int
	for num, _ := range freqMap {
		result = append(result, num)
	}

	// Sorting the result based on frequencies in descending order
	sort.Slice(result, func(i, j int) bool {
		return freqMap[result[i]] > freqMap[result[j]]
	})

	// Returning only the k most frequent elements
	return result[:k]
}
