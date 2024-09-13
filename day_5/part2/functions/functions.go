package functions

func isTriple(a, b, c byte) bool {
	return a == b && b == c
}

func overlapChecker(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if isTriple(s[i], s[i+1], s[i+2]) {
			return false
		}
	}
	return true
}

func pairCounter(s string) bool {
	pairMap := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		// Form the current pair
		pair := string(s[i]) + string(s[i+1])

		// increment the count for this pair
		pairMap[pair]++

		// if a pair appears more than once, return true
		if pairMap[pair] > 1 {
			return true
		}
	}
	return false
}

func newNaughtyOrNice(s string) string {
	if pairCounter(s) && overlapChecker(s) && triples(s) {
		return "nice"
	}
	return "naughty"
}

func triples(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func Counter(str []string) (int, int) {
	var Nccount, Ntcount int

	for _, s := range str {
		state := newNaughtyOrNice(s)
		if state == "naughty" {
			Nccount++
		} else {
			Ntcount++
		}
	}
	return Nccount, Ntcount
}
