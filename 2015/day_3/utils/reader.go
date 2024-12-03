package utils

import (
	"fmt"
	"os"
)

func Reader(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error reading file %q: %v", fileName, err)
	}
	return string(data), nil
}
