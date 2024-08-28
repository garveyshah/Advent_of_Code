package funcs

func SurfaceArea(cup []int) int {
	l, w, h := cup[0], cup[1], cup[2]

	sa := (2*l*w + 2*h*w + 2*h*l)

	if (l == w) || (l == h) || (w == h) {
		sa += 1
	} else {
		sa += 6
	}
	return sa
}

func TotalSa(Jug [][]int) int {
	var totalSa int
	for _, cup := range Jug {
		totalSa += SurfaceArea(cup)
	}
	return totalSa
}
