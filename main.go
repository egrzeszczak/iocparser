package main

import (
	"bufio"
	"fmt"
	"os"
)

// readIoCFromFile reads lines from a file and returns them as a slice.
func readIoCFromFile(filePath string) ([]string, error) {
	// Check if the file exists.
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("File not found: %s", filePath)
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

func main() {
	// Check if the correct number of command-line arguments is provided.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	// Retrieve the file path from the command-line argument.
	filePath := os.Args[1]

	// Read lines from the file using the readIoCFromFile function.
	lines, err := readIoCFromFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the lines or perform further processing as needed.
	fmt.Println("Lines from the file:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
