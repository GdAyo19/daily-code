package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(runLengthEncode(""))
	fmt.Println(runLengthEncode("abcd"))
	fmt.Println(runLengthEncode("zzzzzzzzzzz"))
	fmt.Println(runLengthEncode("aaaaaaaaaaaa"))
	fmt.Println(runLengthEncode("aaabbc"))

}
func runLengthEncode(text string) string {
	result := ""

	if len(text) == 0 {
		return result
	}

	for i := 0; i < len(text); i++ {
		count := 0

		for j := 0; j < len(text); j++ {
			if text[i] == text[j] {
				count++
			}
		}

		if !strings.Contains(result, string(text[i])) {
			result += string(text[i]) + strconv.Itoa(count)
		}
	}

	return result
}
