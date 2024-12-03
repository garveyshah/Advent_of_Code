package main

import (
	"fmt"
	"strconv"

	"2/funcs"
)

func main() {
	data, _ := funcs.Reader("../data.txt")
	var result int
	for _, nums := range data {
		num, _ := strconv.Atoi(nums)
		result += num
	}
	fmt.Println(result)
}
