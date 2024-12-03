package funcs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reader(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, fmt.Errorf("error opening file %v\n: %v", filename, err)
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	data := []string{}

	for fileScan.Scan() {
		data = append(data, fileScan.Text())
	}
	return data, nil
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
