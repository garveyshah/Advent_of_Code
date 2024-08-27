package main

import (
	"fmt"
	"os"
	"strconv"

	"1/brackets"
	"1/utility"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: go run . <FileName>\ngo run . data.txt")
		return
	}

	var FileName string
	var floor uint

	if len(os.Args) == 2 {
		FileName = os.Args[1]
	} else {
		FileName = os.Args[1]
		flo, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		floor = uint(flo)
	}

	content, err := utility.GetData(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(brackets.SortBrackets(string(content)))
	fmt.Println(brackets.GetPosition(content,floor))
}
