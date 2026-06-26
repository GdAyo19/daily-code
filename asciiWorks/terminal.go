package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("sh", "-c", "tput cols 2>/dev/tty")
	cmd.Stdin = os.Stdin
	widths, err := cmd.Output()
	if err != nil {
		return 80
	}
	// var rows, cols int
	// fmt.Sscanf(string(widths), "%d %d", &rows, &cols)

	// row := strings.Fields(string(widths))
	// new, _ := strconv.Atoi(row[1])
	
	width, err := strconv.Atoi(strings.TrimSpace(string(widths)))
	if err != nil {
		return 80
	}
	return width

}
