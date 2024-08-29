package funcs

func HouseTracker(s string) int {
	HouseMap := make(map[[2]int]bool)
	var count int
	House := [2]int{0, 0}

	HouseMap[House] = true
	count++

	for _, seen := range s {
		if seen == '>' {
			House[0] += 1
		} else if seen == '^' {
			House[1] += 1
		} else if seen == 'v' {
			House[1] -= 1
		} else if seen == '<' {
			House[0] -= 1
		}
		if !HouseMap[House] {
			HouseMap[House] = true
			count++
		}
	}
	return count
}
