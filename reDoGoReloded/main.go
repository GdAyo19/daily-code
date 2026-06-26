package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	data, err := os.ReadFile(arg1)
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}

	cxt := string(data)
	words := strings.Fields(cxt)
	words = fixBinAndHex(words)
	words = caseChange(words)

	result := strings.Join(words, " ")
	result = punctuations(result)
	result = checkQuote(result)
	result += "\n"

	err = os.WriteFile(arg2, []byte(result), 0644)
	if err != nil {
		fmt.Println("Unable to write file")
		return
	}

}

func fixBinAndHex(w []string) []string {

	for i := 0; i < len(w); i++ {
		if w[i] == "(hex)" && i > 0 {
			new, _ := strconv.ParseInt(w[i-1], 16, 64)
			w[i-1] = strconv.FormatInt(new, 10)
			w = append(w[:i], w[i+1:]...)
			i--
		}
	}

	return w
}

func caseChange(w []string) []string {

	for i := 0; i < len(w); i++ {
		switch {

		case w[i] == "(up)" && i > 0:
			toUpper(w, i-1, 1)
			w = append(w[:i], w[i+1:]...)
			i--

		case w[i] == "(up" && i+1 < len(w):
			new, _ := strconv.Atoi(strings.TrimSuffix(w[i+1], ")"))
			toUpper(w, i-1, new)
			w = append(w[:i], w[i+2:]...)
			i--
		}

	}

	return w
}

func toUpper(w []string, index, count int) {

	if count <= 0 || index < 0 {
		return
	}

	w[index] = strings.ToUpper(w[index])
	toUpper(w, index-1, count-1)
}

func toLower(w []string, index, count int) {

	if count <= 0 || index < 0 {
		return
	}

	w[index] = strings.ToLower(w[index])
	toLower(w, index-1, count-1)
}

func toCap(w []string, index, count int) {

	if count <= 0 || index < 0 {
		return
	}

	word := []rune(w[index])

	if len(word) > 0 {
		w[index] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
	}

	toUpper(w, index-1, count-1)
}

func punctuations(w string) string {
	s := []rune(w)
	keep := []rune{}
	for i := 0; i < len(s); i++ {
		if IsSign(s[i]) && s[i-1] == ' ' {
			keep = keep[:len(keep)-1]
		}
		if i+1 < len(s) && IsSign(s[i]) && !IsSign(s[i+1]) && s[i+1] != ' ' {
			keep = append(keep, s[i])
			keep = append(keep, ' ')
			continue
		}
		keep = append(keep, s[i])

	}

	return string(keep)
}

func IsSign(w rune) bool {
	switch w {
	case ',', ';', ':', '.', '?', '!':
		return true
	default:
		return false
	}
}

func checkQuote(w string) string {
	s := []rune(w)
	keep := []rune{}
	inQuote := false

	for i := 0; i < len(s); i++ {

		if s[i] == '\'' && i > 0 {
			keep = append(keep, '\'')
			inQuote = !inQuote
			continue
		}
		if inQuote && s[i] == ' ' {
			if s[i-1] == '\'' {
				continue
			}
			if i+1 < len(s) && s[i+1] == '\'' {
				continue
			}
		}
		keep = append(keep, s[i])
	}

	return string(keep)
}

func vowelsFix(w []string) []string {
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(w); i++ {
		if w[i] == "a" || w[i] == "A" {
			next := w[i+1]

			if len(next) > 0 {
				firstWord := next[0]

				if strings.ContainsRune(vowels, rune(firstWord)) {
					if w[i] == "A" {
						w[i] = "An"
					} else {
						w[i] = "an"
					}
				}
			}
		}
	}
	return w
}
