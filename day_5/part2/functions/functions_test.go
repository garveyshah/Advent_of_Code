package functions

import (
	"testing"
)

// Test for Naughty()
func TestNewNaughtyOrNice(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{"Test 0", "qjhvhtzxzqqjkmpb", "nice"},
		{"Test 0.1", "xxyxx", "nice"},
		{"Test 0.2", "uurcxstgmygtbstg", "naughty"},
		{"Test 0.3", "ieodomkazucvgmuy", "naughty"},
		// {"Test 1", "jchzalrnumimnmhp", "naughty"},
		// {"Test 2", "haegwjzuvuyypxyu", "naughty"},
		// {"Test 3", "dvszwmarrgswjxmb", "naughty"},
		// {"Test 4", "ugknbfddgicrmopn", "nice"},
		// {"Test 5", "aaa", "nice"},
		// {"Test 6", "xazegov", "naughty"},
		// {"Test 7", "aei", "naughty"},
		// {"Test 8", "aeiouaeiouaeiou", "naughty"},
		// {"Test 9", "aabbccdd", "naughty"},
		// {"Test 10", "abcdde", "naughty"},
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

func TestPairCounter(t *testing.T) {
	tt := []struct {
		intput string
		want   bool
	}{
		{"aabcaa", true},
		{"aabc", false},
	}

	for _, tc := range tt {
		got := pairCounter(tc.intput)

		if got != tc.want {
			t.Errorf("for input %q Failed - got %v want %v", tc.intput, got, tc.want)
		}
	}
}

func TestOverlapChecker(t *testing.T) {
	tt := []struct {
		input string
		want  bool
	}{
		{"aabbbaa", false},
		{"aabaa", true},
	}

	for _, tc := range tt {
		got := overlapChecker(tc.input)

		if got != tc.want {
			t.Fatalf("for input %q Failed - got %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestTriples(t *testing.T) {
	tt := []struct {
	input string
	want bool
	}{
		{"xyx", true},
		{"aaa", true},
		{"aay",false},
	}

	for _, tc := range tt {
		got := triples(tc.input)

		if got != tc.want {
			t.Errorf("for input %q, got %v, want %v", tc.input, got, tc.want)
		} 
	}
}