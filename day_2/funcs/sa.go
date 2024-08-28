package funcs

func WrappPaper(cup []int) int {
	l, w, h := cup[0], cup[1], cup[2]

	sa := (2*l*w + 2*h*w + 2*h*l)
	m := Min((l * w), (h * w), (h * l))

	sa += m
	return sa
}

func TotalSa(Jug [][]int) int {
	var totalSa int
	for _, cup := range Jug {
		totalSa += WrappPaper(cup)
	}
	return totalSa
}

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
