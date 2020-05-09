package word

import "strings"

// return the number of times a given word appears in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// count returns the number of words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
