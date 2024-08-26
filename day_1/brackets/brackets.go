package brackets

func SortBrackets(s string) uint {
	var move uint
	if s == "" {
		return move
	}
	for _, char := range s {
		switch {
		case char == '(':
			move += 1
		case char == ')':
			move -= 1
		default:
		}
	}

	return move
}
