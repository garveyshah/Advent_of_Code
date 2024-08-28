package funcs
// Ribbon() calcualates the total ribbon quatity
func Ribbon(jug [][]int) int {
	var ribb int
	for _, cup := range jug {
		l := cup[0]
		h := cup[1]
		w := cup[2]

		pr := Min((2*l + 2*h), (2*w + 2*h), (2*l + 2*w))

		ribb += (l * w * h) + pr
	}
	return ribb
}

// WrappPaper() calcultes amount of wrapping paper per present
func WrappPaper(cup []int) int {
	l, w, h := cup[0], cup[1], cup[2]

	sa := (2*l*w + 2*h*w + 2*h*l)
	m := Min((l * w), (h * w), (h * l))

	sa += m
	return sa
}
// TotalSa calculates the total amount of wrapping paper needed for all the gifts
func TotalSa(Jug [][]int) int {
	var totalSa int
	for _, cup := range Jug {
		totalSa += WrappPaper(cup)
	}
	return totalSa
}

// Min() determines the minimum number in a set of three 
func Min(a, b, c int) int {
	min := a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}
	return min
}
