package funcs

import "regexp"

// Counter() counts the number of Nice and Naughty strings
func Counter(data []string) (int, int) {
	var Nice, Naughty int
	for _, str := range data {
		state := NaughtyOrNice(str)
		if state == "nice" {
			Nice++
		} else {
			Naughty++
		}

	}
	return Nice, Naughty
}

// NaughtOrNice() checks if a string is either Nice or Naughty
func NaughtyOrNice(s string) string {
	if nice(s) {
		return "nice"
	}
	return "naughty"
}

// nice() checks if a string has the properties of a Nice string
func nice(s string) bool {
	// Check for at least 3 vowels
	vowels := regexp.MustCompile(`([aeiou].*){3,}`)

	// Ensure it doesn't contain 'ab', 'cd, 'pq', or 'xy'
	forbiddenRegex := regexp.MustCompile(`(ab|cd|pq|xy)`)

	// Perform the checks
	return vowels.MatchString(s) && DoubleLetters(s) && !forbiddenRegex.MatchString(s)
}

// DoubleLetters() checks for the instances of a pairs of double letters in a string
func DoubleLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}
