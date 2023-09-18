package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type IoC struct {
	Value string
	Type  string
	Raw   string
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
	urlPattern := `(https?://[^\s]+)`
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
	domainPattern := `[\w\.-]+\.[a-zA-Z]{2,}`
	domainRegex := regexp.MustCompile(domainPattern)
	emailDomainPattern := `@([\w\.-]+\.[a-zA-Z]{2,})`
	emailDomainRegex := regexp.MustCompile(emailDomainPattern)

	// Loop through lines and check for IoCs.
	for _, line := range lines {
		if urlRegex.MatchString(line) {
			// Detected URL
			urlMatches := urlRegex.FindStringSubmatch(line)
			if len(urlMatches) > 1 {
				url := urlMatches[1]
				iocs = append(iocs, IoC{Value: url, Type: "URL", Raw: line})
			}

			// Extract IPv4 addresses from the URL
			ipMatches := ipRegex.FindAllString(line, -1)
			if len(ipMatches) > 0 {
				for _, ip := range ipMatches {
					iocs = append(iocs, IoC{Value: ip, Type: "IPv4", Raw: line})
				}
			} else {
				// Extract Domains from the URL
				domainMatch := domainRegex.FindAllString(line, -1)
				for range domainMatch {
					iocs = append(iocs, IoC{Value: domainMatch[0], Type: "Domain", Raw: line})
				}
			}

		} else if ipRegex.MatchString(line) {
			// Detected IPv4 Address
			iocs = append(iocs, IoC{Value: line, Type: "IPv4", Raw: line})
		} else if sha256Regex.MatchString(line) {
			// Detected SHA256 Hash
			iocs = append(iocs, IoC{Value: line, Type: "SHA256", Raw: line})
		} else if sha1Regex.MatchString(line) {
			// Detected SHA1 Hash
			iocs = append(iocs, IoC{Value: line, Type: "SHA1", Raw: line})
		} else if md5Regex.MatchString(line) {
			// Detected MD5 Hash
			iocs = append(iocs, IoC{Value: line, Type: "MD5", Raw: line})
		} else if emailRegex.MatchString(line) {
			// Detected Email Address
			iocs = append(iocs, IoC{Value: line, Type: "Email", Raw: line})

			// Extract Domains from the URL
			domainMatch := emailDomainRegex.FindStringSubmatch(line)
			if len(domainMatch) > 1 {
				iocs = append(iocs, IoC{Value: domainMatch[1], Type: "Domain", Raw: line})
			}

		} else if domainRegex.MatchString(line) {
			// Detected domain
			iocs = append(iocs, IoC{Value: line, Type: "Domain", Raw: line})
		} else {
			// It's an Unknown type
			iocs = append(iocs, IoC{Value: line, Type: "Unknown", Raw: line})
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

	// Create a map to store grouped and unique IoCs by Type.
	groupedUniqueIoCs := make(map[string]map[string]struct{})

	// Detect IoC types and group them by Type while ensuring uniqueness.
	iocs := detectIoCTypes(lines)
	var typeOrder []string

	for _, ioc := range iocs {
		// Create a map for each Type if it doesn't exist.
		if _, exists := groupedUniqueIoCs[ioc.Type]; !exists {
			groupedUniqueIoCs[ioc.Type] = make(map[string]struct{})
			typeOrder = append(typeOrder, ioc.Type) // Add the detected Type to typeOrder.
		}

		// Add the IoC Value to the map to ensure uniqueness.
		groupedUniqueIoCs[ioc.Type][ioc.Value] = struct{}{}
	}

	// Sort the typeOrder slice alphabetically.
	sort.Strings(typeOrder)

	// Print the grouped, sorted, and unique IoCs.
	for _, t := range typeOrder {
		if ioCsOfType, exists := groupedUniqueIoCs[t]; exists {
			// Collect the IoCs for this type.
			var sortedIoCs []string
			for ioc := range ioCsOfType {
				sortedIoCs = append(sortedIoCs, ioc)
			}

			// Sort the IoCs alphabetically.
			sort.Strings(sortedIoCs)

			// Print the IoCs.
			fmt.Printf("%s:\n", t)
			for _, ioc := range sortedIoCs {
				fmt.Printf("%s\n", ioc)
			}
			fmt.Println() // Add a blank line between groups.
		}
	}
}
