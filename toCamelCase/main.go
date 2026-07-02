package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("hello world")) // helloWorld
	fmt.Println(ToCamelCase("hello_world")) // helloWorld
	fmt.Println(ToCamelCase("hello-world")) // helloWorld
	fmt.Println(ToCamelCase("GO LANGUAGE")) // goLanguage
}

func ToCamelCase(str string) string {

	// str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.ReplaceAll(str, "-", " ")
	words := strings.Fields(str)

	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])

	for _, word := range words[1:] {
		word = strings.ToLower(word)
		result += strings.ToUpper(word[:1]) + word[1:]
	}

	return result
}
