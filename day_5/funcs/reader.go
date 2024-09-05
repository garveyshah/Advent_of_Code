package funcs

import (
	"bufio"
	"fmt"
	"os"
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
