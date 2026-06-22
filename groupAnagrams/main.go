package main

import "fmt"

func main() {
	// words := []string{"eat", "tea", "bat"}
	fmt.Println(groupAnagrams([]string{"abc", "def"}))
}

func groupAnagrams(words []string) [][]string {

	groups := make(map[string][]string)

	for _, word := range words {
		k := key(word)
		groups[k] = append(groups[k], word)
	}

	result := [][]string{}

	for _, group := range groups {
		result = append(result, group)

	}

	return result
}

func key(word string) string {
	r := []rune(word)

	for i := 0; i < len(r); i++ {
		for j := 0; j < i; j++ {
			if r[i] < r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)

}
