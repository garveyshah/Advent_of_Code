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

	data, err := funcs.Reader("../part1/test_data/" + os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	Nice, Naughty := functions.Counter(data)

	fmt.Println("Nice - ", Nice, "\n", "Naughty - ", Naughty)
}
