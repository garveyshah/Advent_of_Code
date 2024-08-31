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
	var count int

	// Starting positions for Santa and the robot
	santa := [2]int{0, 0}
	robot := [2]int{0, 0}

	// Mark the inital position as visited
	HouseMap[santa] = true
	count++

	for i, seen := range s {
		// Determine whether it;s santa's or robot's turn to
		var current *[2]int
		if i%2 == 0 {
			current = &santa
		} else {
			current = &robot
		}

		if seen == '>' {
			current[0] += 1
		} else if seen == '^' {
			current[1] += 1
		} else if seen == 'v' {
			current[1] -= 1
		} else if seen == '<' {
			current[0] -= 1
		}

		// Check if this position has been visited before
		if !HouseMap[*current] {
			HouseMap[*current] = true
			count++
		}
	}
	return count
}
