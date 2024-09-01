package main

import (
	"fmt"
	"os"

	"day4/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <string>\nExample : go run . \"abcdef\"")
		return
	}

	code := funcs.MD5(os.Args[1])
	fmt.Println("The answer for " + os.Args[1] + " is : " + code)
}
