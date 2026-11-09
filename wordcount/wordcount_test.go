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
		want int
	}{
		{"when text is empty, it should return nothing", "", 3, 0},
		{"when n is zero, it should return nothing", "a bb", 0, 0},
		{"when n is negative, it should return nothing", "a bb", -1, 0},
		{"when n exceeds the vocabulary, it should return every word", "a bb", 10, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := len(Longest(tc.text, tc.n)); got != tc.want {
				t.Errorf("len(Longest(%q, %d)) = %d, want %d", tc.text, tc.n, got, tc.want)
			}
		})
	}

	got := Longest("éé aaa bb aa", 4)
	want := []string{"aaa", "aa", "bb", "éé"}
	if len(got) != len(want) {
		t.Fatalf("Longest returned %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Longest returned %v, want %v", got, want)
			break
		}
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
