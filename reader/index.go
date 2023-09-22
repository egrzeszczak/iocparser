package reader

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filePath string) ([]string, error) {
	// Check if the file exists.
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to store the lines from the file.
	var lines []string

	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
