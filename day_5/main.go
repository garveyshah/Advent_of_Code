package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")
	}

	result := funcs.StringSort(os.Args[1])
	fmt.Println(result)
}
