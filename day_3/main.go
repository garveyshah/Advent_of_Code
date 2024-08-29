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
	gifted2 := funcs.HouseTracker2(data)
	fmt.Println("Houses In Year \"1\" gifted are : ", gifted)
	fmt.Println("Houses In Year \"2\" gifted are : ", gifted2)
}
