package main

import (
	"reflect"
	"sort"
)

func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sChars := []byte(s)
	tChars := []byte(t)

	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] < sChars[j]
	})

	sort.Slice(tChars, func(i, j int) bool {
		return tChars[i] < tChars[j]
	})

	if reflect.DeepEqual(sChars, tChars) {
		return true
	}

	return false
}
