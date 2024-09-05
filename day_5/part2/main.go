package main

import (
	"fmt"
	"os"

	"day5/part1/funcs"
)

func main() {
	if len(os.Args[1]) != 2 {
		fmt.Println("Usage : go run . <filename>\nExample : go run . data.txt")
		return
	}

	data, err := funcs.Reader(os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	Nice, Naughty := functions.Counter(data)

	fmt.Println("Nice - ", Nice, "\n", "Naughty - ", Naughty)
}
