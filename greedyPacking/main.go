package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wrapText([]string{"This", "is", "an", "example", "text"}, 10))
}

func wrapText(words []string, maxWidth int) []string {
	text := []string{}
	currentLine := []string{}
	currentLength := 0

	for _, word := range words {

		if len(currentLine) == 0 {
			currentLine = append(currentLine, word)
			currentLength = len(word)
		} else if (currentLength + 1 + len(word)) <= maxWidth {
			currentLine = append(currentLine, word)
			currentLength += 1 + len(word)
		} else {
			lines := strings.Join(currentLine, " ")
			padding := maxWidth - len(lines)
			text = append(text, lines+strings.Repeat(" ", padding))
			currentLine = []string{word}
			currentLength = len(word)
		}

	}
	if len(currentLine) > 0 {
		lines := strings.Join(currentLine, " ")
		padding := maxWidth - len(lines)
		text = append(text, lines+strings.Repeat(" ", padding))

	}

	return text
}
