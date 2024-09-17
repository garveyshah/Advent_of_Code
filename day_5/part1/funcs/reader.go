package funcs

import (
	"fmt"
	"os"
	"strings"
)

func Reader(filename string) ([]string, error) {
	// file, err := os.Open(filename)
	// if err != nil {
	// 	return []string{}, fmt.Errorf("error opening file %v\n: %v", filename, err)
	// }
	// defer file.Close()

	// fileScan := bufio.NewScanner(file)
	// fileScan.Split(bufio.ScanLines)
	// data := []string{}

	// for fileScan.Scan() {
	// 	data = append(data, fileScan.Text())
	// }
	// return data, nil

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %q: %v ", filename, err)
	}

	if len(file) < 1 {
		return nil, fmt.Errorf("file is empty %q", filename)

	}

	str := string(file)
	words := strings.Fields(str)
	return words, nil
}
