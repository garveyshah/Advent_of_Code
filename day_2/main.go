package main

import "fmt"

func main() {
	l, w, h := 2, 3, 4
	l1, w2, h3 := 1, 1, 10

	TFeet := SurfaceArea(l, w, h)
	TFeet2 := SurfaceArea(l1, w2, h3)
	fmt.Println("instance A:= ", TFeet)
	fmt.Println("Instance B := ", TFeet2)
}

func SurfaceArea(l, w, h int) int {
	sa := (2*l*w + 2*h*w + 2*h*l)

	if (l == w) || (l == h) || (w == h) {
		sa += 1
	} else {
		sa += 6
	}
	return sa
}
