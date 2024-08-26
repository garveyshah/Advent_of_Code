package main

import (
	"fmt"
	"os"

	"1/brackets"
	"1/utility"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <FileName>\ngo run . data.txt")
		return
	}

	FileName := os.Args[1]
	content, err := utility.GetData(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(brackets.SortBrackets(string(content)))
}

