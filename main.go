package main

import (
	"fmt"

	detect "github.com/egrzeszczak/iocparser/detect"
	"github.com/egrzeszczak/iocparser/enhancer"
	"github.com/egrzeszczak/iocparser/input"
	"github.com/egrzeszczak/iocparser/reader"
)

func main() {

	// 1. Read input args
	filePath, _, _ := input.GetArgs()

	// 2. Read file given in args
	fileLines, err := reader.Read(filePath)
	if err != nil {
		fmt.Errorf("Error while reading file: %s", err)
		return
	}

	// 3. Detect IoCs in each line
	fileIOCs := detect.Detect(fileLines)
	fmt.Printf("%v", fileIOCs)

	// 4. Create a link to each IoC to a reputation service
	for _, ioc := range fileIOCs {
		fmt.Printf("%v", enhancer.CreateLinks(ioc))
	}

	// 5. Display
	// WORK IN PROGRESS

	// // Read lines from the file using the readIoCsFromFile function.
	// lines, err := readIoCsFromFile(filePath)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Detect IoC types and group them by Type while ensuring uniqueness.
	// iocs := detectIoCs(lines)

	// // Print IoCs to Markdown style
	// printIoCsToMarkdown(iocs, config)
}
