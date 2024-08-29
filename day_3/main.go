package main

import (
	"fmt"
	"os"

	"day3/funcs"
	"day3/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : go run . <filename>\nExample: go run . data.txt")
		return
	}
	data, err := utils.Reader(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	gifted := funcs.HouseTracker(data)
	fmt.Println("Houses gifted are : ", gifted)
}
