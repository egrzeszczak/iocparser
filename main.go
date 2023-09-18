package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type IoC struct {
	Value string
	Type  string
}

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

// detectIoCTypes is a function for detecting IoC types and returns them as a slice of IoC structs.
func detectIoCTypes(lines []string) []IoC {
	var iocs []IoC

	// Define regular expressions to match URLs, IPv4 addresses, domains, SHA1, SHA256, MD5, and email addresses.
	urlPattern := `https?://[^\s]+`
	urlRegex := regexp.MustCompile(urlPattern)
	ipPattern := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`
	ipRegex := regexp.MustCompile(ipPattern)
	sha1Pattern := `[0-9a-fA-F]{40}`
	sha1Regex := regexp.MustCompile(sha1Pattern)
	sha256Pattern := `[0-9a-fA-F]{64}`
	sha256Regex := regexp.MustCompile(sha256Pattern)
	md5Pattern := `[0-9a-fA-F]{32}`
	md5Regex := regexp.MustCompile(md5Pattern)
	emailPattern := `[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,4}`
	emailRegex := regexp.MustCompile(emailPattern)
	domainPattern := `[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	domainRegex := regexp.MustCompile(domainPattern)

	// Loop through lines and check for IoCs.
	for _, line := range lines {
		if urlRegex.MatchString(line) {
			// Detected URL
			iocs = append(iocs, IoC{Value: line, Type: "URL"})
		}
	}

	for _, line := range lines {
		if ipRegex.MatchString(line) {
			// Detected IPv4 Address
			iocs = append(iocs, IoC{Value: line, Type: "IPv4"})
		} else if sha1Regex.MatchString(line) {
			// Detected SHA1 Hash
			iocs = append(iocs, IoC{Value: line, Type: "SHA1"})
		} else if sha256Regex.MatchString(line) {
			// Detected SHA256 Hash
			iocs = append(iocs, IoC{Value: line, Type: "SHA256"})
		} else if md5Regex.MatchString(line) {
			// Detected MD5 Hash
			iocs = append(iocs, IoC{Value: line, Type: "MD5"})
		} else if emailRegex.MatchString(line) {
			// Detected Email Address
			iocs = append(iocs, IoC{Value: line, Type: "Email"})
		} else if domainRegex.MatchString(line) {
			// Detected domain
			iocs = append(iocs, IoC{Value: line, Type: "Domain"})
		} else {
			// It's an Unknown type
			iocs = append(iocs, IoC{Value: line, Type: "Unknown"})
		}
	}

	return iocs
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
	// fmt.Println("Lines from the file:")
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	iocs := detectIoCTypes(lines)

	for _, ioc := range iocs {
		fmt.Printf("Type: %s, Value: %s\n", ioc.Type, ioc.Value)
	}

	// fmt.Printf("%v", detectIoCTypes(lines))
}
