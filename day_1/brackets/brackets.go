package brackets

func SortBrackets(s string) uint {
	var floor uint
	if s == "" {
		return floor
	}
	for _, char := range s {
		switch {
		case char == '(':
			floor += 1
		case char == ')':
			floor -= 1
		default:
		}
	}

	return floor
}

func GetPosition(s string, floor uint) int {
	var move uint

		for i, char := range s {
			switch {
			case char == '(':
				move += 1
			case char == ')':
				move -= 1
			
			}
			if move == floor {
				return i+1
			}
		}
	return -1
}
