package main

import (
	// "fmt"
	"strings"
)

func fixAlignRight(line string, width int) string {
	spaces := width - len(line)

	if spaces < 0 {
		spaces = 0
	}

	return strings.Repeat(" ", spaces) + line
}

func fixAlignCenter(line string, width int) string {
	spaces := (width - len(line)) / 2
	// print(len(line))

	// println(spaces)

	if spaces < 0 {
		spaces = 0
	}

	return strings.Repeat(" ", spaces) + line
}

func fixAlignJustify(line string, width int) string {

	if line == "" {
		return ""
	}

	spaces := width - len(line)

	if spaces <= 0 {
		return line
	}

	// add spaces evenly between existing characters
	gaps := len(line) - 1

	if gaps == 0 {
		return line
	}

	spacePerGap := spaces / gaps
	extra := spaces % gaps

	result := ""

	for i, ch := range line {

		result += string(ch)

		if i < gaps {

			result += strings.Repeat(" ", spacePerGap)

			if extra > 0 {
				result += " "
				extra--
			}
		}
	}

	return result
}
