package main

import (
	"fmt"
	"os"

	"2/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filePath>\nexample: go run . data.txt")
		return
	}

	filename := os.Args[1]

	data, err := funcs.Reader(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	Jug := funcs.DecodeData(data)
	fmt.Println(Jug)
	totalSa := funcs.TotalSa(Jug)
	fmt.Println("Total SA ==", totalSa)
	// l, w, h := 2, 3, 4
	// l1, w2, h3 := 1, 1, 10

	// TFeet := funcs.SurfaceArea(l, w, h)
	// TFeet2 := funcs.SurfaceArea(l1, w2, h3)
	// fmt.Println("instance A:= ", TFeet)
	// fmt.Println("Instance B := ", TFeet2)
}
