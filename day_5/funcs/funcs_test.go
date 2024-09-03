package funcs

import (
	"testing"
)

// Test for Naughty()
func TestNaughtyOrNice(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{"Test 1", "jchzalrnumimnmhp", "naughty"},
		{"Test 2", "haegwjzuvuyypxyu", "naughty"},
		{"Test 3", "dvszwmarrgswjxmb", "naughty"},
		{"Test 4", "ugknbfddgicrmopn", "nice"},
		{"Test 5", "aaa", "nice"},
		{"Test 6", "xazegov", "nice"},
		{"Test 7", "aei", "nice"},
		{"Test 8", "aeiouaeiouaeiou", "nice"},
		{"Test 9", "aabbccdd", "nice"},
		{"Test 9", "abcdde", "nice"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := NaughtyOrNice(tc.input)
			if tc.want != got {
				t.Errorf("%q, Failed: got %s, want %s", tc.name, got, tc.want)
			}
		})
	}
}
