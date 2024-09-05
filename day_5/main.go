package main

import (
	"fmt"
	"os"

	"day5/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")
		return
	}

	data, err := funcs.Reader(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}


	Nice, Naughty := funcs.Counter(data)

	fmt.Println("Nice - ", Nice)
	fmt.Println("Naughty - ", Naughty)
}
