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

func HouseTracker2(s string) int {
		HouseMap := make(map[[2]int]bool)
		HouseMap1 := make(map[[2]int]bool)

	var count,count1 int
	House := [2]int{0, 0}

	HouseMap[House] = true
	count++
	HouseMap1[House] = true
	count1++

	for i, seen := range s {

		if i%2 == 0 {
			if seen == '>' {
				House[0] += 1
			} else if seen == '^' {
				House[1] += 1
			} else if seen == 'v' {
				House[1] -= 1
			} else if seen == '<' {
				House[0] -= 1
			}
			if !HouseMap1[House] {
				HouseMap1[House] = true
				count1++
			}
			//count1 -= 1
		} else {
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
	}
	tcount := count + count1
	return tcount
}
