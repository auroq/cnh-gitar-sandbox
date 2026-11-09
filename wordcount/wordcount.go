package wordcount

import (
	"sort"
	"strings"
	"unicode"
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
