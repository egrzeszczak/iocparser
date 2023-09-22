package input

import (
	"flag"
	"fmt"
	"os"
)

func GetArgs() (string, string, string) {
	// Define command-line flags
	format := flag.String("f", "", "Format (optional)")
	outputFile := flag.String("o", "", "Output file (optional)")
	inputFile := flag.String("i", "", "Input file (required)")

	// Parse the command-line arguments
	flag.Parse()

	// Check if the -i flag (input) is provided and not empty
	if *inputFile == "" {
		fmt.Println("Error: Input file (-i) is required")
		flag.Usage()
		os.Exit(1)
	}

	// Your application logic can go here
	fmt.Printf("Input file: %s\n", *inputFile)
	if *outputFile != "" {
		fmt.Printf("Output file: %s\n", *outputFile)
	}
	if *format != "" {
		fmt.Printf("Format: %s\n", *format)
	}

	return *inputFile, *outputFile, *format
}
