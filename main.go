package main

import (
	"os"
	"sort"
)

type WordCount struct {
	word  []byte
	count int
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	words := extractWords(data)
	counts := countWords(words)

	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count != counts[j].count {
			return counts[i].count > counts[j].count
		}

		return bytesLess(counts[j].word, counts[i].word)
	})

	limit := 20
	if len(counts) < 20 {
		limit = len(counts)
	}

	for i := 0; i < limit; i++ {
		writeOutput(counts[i].count, counts[i].word)
	}
}

func bytesLess(a, b []byte) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return true
		}

		if a[i] > b[i] {
			return false
		}
	}

	return len(a) < len(b)
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}

	return b
}

func extractWords(data []byte) [][]byte {
	var words [][]byte
	var current []byte

	for _, b := range data {
		if isLetter(b) {
			current = append(current, toLower(b))
		} else {
			if len(current) > 0 {
				words = append(words, current)
				current = nil
			}
		}
	}

	if len(current) > 0 {
		words = append(words, current)
	}

	return words
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func countWords(words [][]byte) []WordCount {
	var counts []WordCount

	for _, word := range words {
		found := false
		for i := range counts {
			if bytesEqual(counts[i].word, word) {
				counts[i].count++
				found = true
				break
			}
		}

		if !found {
			counts = append(counts, WordCount{
				word:  word,
				count: 1,
			})
		}
	}

	return counts
}

func writeOutput(count int, word []byte) {
	var digits []byte
	if count == 0 {
		digits = []byte{'0'}
	} else {
		for count > 0 {
			digits = append([]byte{byte('0' + count%10)}, digits...)
			count /= 10
		}
	}

	var line []byte
	padding := 7 - len(digits)

	for i := 0; i < padding; i++ {
		line = append(line, ' ')
	}

	line = append(line, digits...)
	line = append(line, ' ')
	line = append(line, word...)
	line = append(line, '\n')

	os.Stdout.Write(line)
}
