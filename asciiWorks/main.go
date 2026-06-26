package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	var color string
	var text string
	var align string
	var substring string

	for _, arg := range args {

		switch {
		case strings.HasPrefix(arg, "--color="):
			color = strings.TrimPrefix(arg, "--color=")
		case strings.HasPrefix(arg, "--align="):
			align = strings.TrimPrefix(arg, "--align=")
		default:
			if text == "" {
				text = arg
			} else {
				substring = arg
			}
		}
	}

	result := AsciiColor(text, color, substring)

	switch align {
	case "right":
		result = asciiAlign(result, fixAlignRight)
	case "center":
		result = asciiAlign(result, fixAlignCenter)
	case "justify":
		result = asciiAlign(result, fixAlignJustify)
	case "left", "":
		// no alignment transformation needed
	default:
		fmt.Fprintf(os.Stderr, "Invalid alignment option %q. Use 'left', 'right', 'center', or 'justify'.\n", align)
		os.Exit(1)
	}

	fmt.Print(result)
}

func asciiAlign(text string, fn func(string, int) string) string {
	width := GetTerminalWidth()

	lines := strings.Split(text, "\n")
	result := []string{}

	for _, line := range lines {
		result = append(result, fn(line, width))
	}

	return strings.Join(result, "\n")
}