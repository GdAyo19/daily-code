package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [SUBSTRING] [STRING]\nOR YOU USE go run . [OPTION] [STRING]")
		return
	}

	arg1 := os.Args[1] // for color flag
	var substring string
	var text string
	var colorFlag string

	if len(os.Args) == 4 {
		substring = os.Args[2]
		text = os.Args[3]
	} else {
		text = os.Args[2]
	}

	if strings.HasPrefix(arg1, "--color=") {
		colorFlag = strings.TrimPrefix(arg1, "--color=")
	} else {
		fmt.Print("Please start flag with '--color='")
	}

	fmt.Println(AsciiColor(text, colorFlag, substring))

}

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

				start := strings.Index(word, target)

				if start != -1 {
					if r >= start && r < start+len(target) {
						match = true
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
