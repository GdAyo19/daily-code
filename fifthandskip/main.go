package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {

	strs := []rune(strings.ReplaceAll(str, " ", ""))

	if len(strs) < 5 {
		return "Invalid input\n"
	}
	if len(strs) == 0{ 
		return "\n"
	}


	box := []rune{}
	
	count := 0

	for i := range strs {
		if count == 5 {
			box = append(box, ' ')
			count = 0
			continue
		}
		box = append(box, strs[i])
		count++

	}
	return string(box)
}
