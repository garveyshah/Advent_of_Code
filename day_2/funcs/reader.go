package funcs

import (
	"strconv"
	"strings"
)

func Reader(FileName string) []int {
	return []int{}
}

func DecodeData(s []string) [][]int {
	var Jug [][]int
	var cup []int
	for _, str := range s {
		Nstr := strings.Split(str, "x")

		for _, char := range Nstr {
			num, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			cup = append(cup, num)
		}
		Jug = append(Jug, cup)
		cup = []int{}
	}
	return Jug
}
