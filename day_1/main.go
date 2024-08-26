package main

import (
	"fmt"
	"os"

	"1/brackets"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <FileName>\ngo run . data.txt")
		return
	}

	FileName := os.Args[1]

	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("Error accessing file: ", err)
		return
	}
	defer file.Close()

	content, err := os.ReadFile(FileName)
	if err != nil {
		fmt.Println("Error Reading file: ", err)
		return
	}

	fmt.Println(brackets.SortBrackets(string(content)))
}
