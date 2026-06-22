package main

import "fmt"

func main() {
	fmt.Println(findAllOccurrences("aaaaaaa", "aaa"))
	fmt.Println(findAllOccurrences("abcabcabc", "abc"))
	fmt.Println(findAllOccurrences("mississippi", "issi"))
	fmt.Println(findAllOccurrences("abc", ""))
	fmt.Println(findAllOccurrences("", "a"))

}

func findAllOccurrences(text, pattern string) []int {
	result := []int{}

	if pattern == "" || len(pattern) > len(text) {
		return result
	}

	for i := 0; i <= len(text)-len(pattern); i++ {
		match := true

		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			result = append(result, i)
		}
	}

	return result
}
