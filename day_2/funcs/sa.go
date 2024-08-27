package funcs

func SurfaceArea(l, w, h int) int {
	sa := (2*l*w + 2*h*w + 2*h*l)

	if (l == w) || (l == h) || (w == h) {
		sa += 1
	} else {
		sa += 6
	}
	return sa
}
