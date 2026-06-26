package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiColor(text, color, substring string) string {

	reset := "\033[0m"

	if color == "" {
		fmt.Println("Please input a valid color or input a color if not there")
		return ""
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Unable to read file")
		return ""
	}

	cnx := strings.Split(string(data), "\n")
	sentence := strings.ReplaceAll(string(text), "\r\n", "\n")
	words := strings.Split(sentence, "\n")
	result := ""

	for _, word := range words {

		if word == "" {
			result += "\n"
			continue
		}

		target := substring
		if substring == "" {
			target = word
		}

		for i := 1; i <= 8; i++ {

			for r, ch := range word {
				if ch < 32 || ch > 126 {
					continue
				}
				index := ((int(ch) - 32) * 9)

				if index+r >= len(cnx) {
					continue
				}

				ascii := cnx[index+i]

				match := false

				for start := 0; start+len(target) <= len(word); start++ {
					if word[start:start+len(target)] == target {
						if r >= start && r < start+len(target) {
							match = true
						}
					}
				}

				if match {
					result += getColor(color)
					result += ascii
					result += reset
				} else {
					result += ascii
				}
			}
			result += "\n"
		}

	}

	return result
}

func getColor(name string) string {

	color, ok := colors[name]
	if ok {

		return color
	}
	return ""

}

var colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[38;2;255;165;0m",
}
