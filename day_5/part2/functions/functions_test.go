package functions

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
		{"Test 6", "xazegov", "naughty"},
		{"Test 7", "aei", "naughty"},
		{"Test 8", "aeiouaeiouaeiou", "naughty"},
		{"Test 9", "aabbccdd", "naughty"},
		{"Test 10", "abcdde", "naughty"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := newNaughtyOrNice(tc.input)
			if tc.want != got {
				t.Errorf("%q, Failed: got %s, want %s", tc.name, got, tc.want)
			}
		})
	}
}
