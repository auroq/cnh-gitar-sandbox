package wordcount

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name string
		text string
		word string
		want int
	}{
		{"when text is empty, it should count zero", "", "go", 0},
		{"when a word repeats with mixed case, it should count it once per occurrence", "Go go GO", "go", 3},
		{"when words are split by punctuation, it should still count them", "go,go;go.", "go", 3},
		{"when a word has an apostrophe, it should keep it as one word", "don't don't", "don't", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.text)[tc.word]; got != tc.want {
				t.Errorf("Count(%q)[%q] = %d, want %d", tc.text, tc.word, got, tc.want)
			}
		})
	}
}

func TestLongest(t *testing.T) {
	cases := []struct {
		name string
		text string
		n    int
		want []string
	}{
		{"when text is empty, it should return nothing", "", 3, []string{}},
		{"when n is zero, it should return nothing", "a bb", 0, []string{}},
		{"when n is negative, it should return nothing", "a bb", -1, []string{}},
		{"when n exceeds the vocabulary, it should return every word longest first", "a bb", 10, []string{"bb", "a"}},
		{"when n is smaller than the vocabulary, it should return the n longest words", "a bbb cc", 2, []string{"bbb", "cc"}},
		{"when words tie on length, it should order them alphabetically", "aaa aa bb", 3, []string{"aaa", "aa", "bb"}},
		{"when words are repeated, it should return each word once", "aaa aaa bb", 5, []string{"aaa", "bb"}},
		{"when words have multibyte runes, it should measure length in runes", "éé aaa bb aa", 4, []string{"aaa", "aa", "bb", "éé"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Longest(tc.text, tc.n)
			if len(got) != len(tc.want) {
				t.Fatalf("Longest(%q, %d) = %v, want %v", tc.text, tc.n, got, tc.want)
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Fatalf("Longest(%q, %d) = %v, want %v", tc.text, tc.n, got, tc.want)
				}
			}
		})
	}
}

func TestTop(t *testing.T) {
	cases := []struct {
		name string
		text string
		n    int
		want int
	}{
		{"when n is larger than the vocabulary, it should return every word", "a b c", 10, 3},
		{"when n is smaller than the vocabulary, it should return n words", "a b c", 2, 2},
		{"when n is zero, it should return nothing", "a b c", 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := len(Top(tc.text, tc.n)); got != tc.want {
				t.Errorf("len(Top(%q, %d)) = %d, want %d", tc.text, tc.n, got, tc.want)
			}
		})
	}
}
