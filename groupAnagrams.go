package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// Convert the string to a sorted array of characters
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		sortedStr := string(strBytes)

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
