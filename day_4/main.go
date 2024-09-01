package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <string>\nExample : go run . \"abcdef\"")
		return
	}
}
