package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	words := []string{
		"hello",
		"world",
		"make",
		"a",
		"way",
		"(up,3)",
		"okay",
	}
	x := removeWord([]string{"hello", "my", "name", "is", "john", "(up)"}, "my")
	fmt.Println(applyUpper(x))
	fmt.Println(applyUpper(words))
}
func removeWord(words []string, target string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		if words[i] != target {
			result = append(result, words[i])
		}
	}

	return result
}

func applyUpper(words []string) []string {

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-2] = strings.ToUpper(words[i-2])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(up,") {
			count, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(words[i], "(up,"), ")"))
			for j := 1; j <= count; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

// func cleanText(s string) string {

// }
