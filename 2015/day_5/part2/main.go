package main

import (
	"fmt"
	"os"

	"day5/part1/funcs"
	"day5/part2/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <filename>\nExample : go run . data.txt")
		return
	}

	data, err := funcs.Reader(os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	var count, val int
	for _, word := range data {
		if functions.IsNice(word) {
			count++
		} else {
			val++
		}
	}

	fmt.Printf("Nice - %d\nNaughty - %d\n", count, val)
}
