package functions

// Counter() counts the number of Nice and Naughty strings
func Counter(data []string) (int, int) {
	var Nice, Naughty int
	for _, str := range data {
		state := newNaughtyOrNice(str)
		if state == "nice" {
			Nice++
		} else {
			Naughty++
		}

	}
	return Nice, Naughty
}

// newNaughtOrNice() checks if a string is either Nice or Naughty
func newNaughtyOrNice(s string) string {
	if modifiedNice(s) {
		return "nice"
	}
	return "naughty"
}

// modifiedNice() checks if a string has the properties of a Nice string
func modifiedNice(s string) bool {
	Wmap := make(map[string]int)
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == s[i+1] && s[i] != s[i+2] && i < len(s)-2 {
			Wmap[string(s[i])+string(s[i+1])]++
		} else if i == len(s)-2 && s[i] == s[i+1] {
			Wmap[string(s[i])+string(s[i+1])]++
		}
		if s[i] == s[i+2] && i < len(s)-2{
			return true
		}
	}

	for _, count := range Wmap {
		if count%2 != 0 {
			return false
		}
	}

	return true
}
