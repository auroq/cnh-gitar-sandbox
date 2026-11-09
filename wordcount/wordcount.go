package wordcount

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Entry struct {
	Word  string
	Count int
}

func Count(text string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.FieldsFunc(strings.ToLower(text), isSeparator) {
		counts[w]++
	}
	return counts
}

func Top(text string, n int) []Entry {
	counts := Count(text)
	entries := make([]Entry, 0, len(counts))
	for w, c := range counts {
		entries = append(entries, Entry{Word: w, Count: c})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Word < entries[j].Word
	})
	if n < len(entries) {
		entries = entries[:n]
	}
	return entries
}

func isSeparator(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
}

func Longest(text string, n int) []string {
	words := make([]string, 0)
	for w := range Count(text) {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		li, lj := utf8.RuneCountInString(words[i]), utf8.RuneCountInString(words[j])
		if li != lj {
			return li > lj
		}
		return words[i] < words[j]
	})
	if n < 0 {
		n = 0
	}
	if n > len(words) {
		n = len(words)
	}
	return words[:n]
}
