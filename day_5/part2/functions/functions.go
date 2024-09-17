package functions

func overlapChecker(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
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

func IsNice(s string) bool {
	return  pairCounter(s) && overlapChecker(s) && intervaledPair(s)
}

func intervaledPair(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

