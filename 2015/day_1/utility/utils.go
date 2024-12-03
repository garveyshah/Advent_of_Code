package utility

import (
	"fmt"
	"os"
)

func GetData(FileName string) (string, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return "", fmt.Errorf("error accessing file: %v", err)
	}
	defer file.Close()

	content, err := os.ReadFile(FileName)
	if err != nil {
		return "", fmt.Errorf("error Reading file: %v", err)
	}
	return string(content), nil
}
