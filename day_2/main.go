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
	totalSa := funcs.TotalSa(Jug)
	ribbon := funcs.Ribbon(Jug)
	fmt.Println("Total Wrapping paper ==", totalSa)
	fmt.Println("Total Ribbon == ", ribbon)
}
