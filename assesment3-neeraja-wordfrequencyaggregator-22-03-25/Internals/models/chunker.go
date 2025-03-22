package models

import (
	"bufio"
	"strings"
	"unicode"
)

func SplitWords(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		word := scanner.Text()

		words = append(words, strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r)
		}))
	}
	return words
}
