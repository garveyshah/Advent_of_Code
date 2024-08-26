package brackets

func SortBrackets(s string) uint {
	var move uint
	for _, char := range s {
		if char == '(' {
			move += 1
		} else if move == ')' {
			move -= 1
		}
	}
	return move
}
