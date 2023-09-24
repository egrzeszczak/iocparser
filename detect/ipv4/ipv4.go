package ipv4

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
	"github.com/jakewarren/defang"
)

// IPv4 represents a parsed IPv4 address with various formats.
type IPv4 struct {
	Fanged   string
	Defanged string
	Encoded  string
	Original string
}

var IPv4Regex *regexp2.Regexp

func init() {
	var err error
	// IPv4Regex is initialized with a regular expression pattern for detecting public IPv4 addresses.
	IPv4Regex, err = regexp2.Compile(`(?!(192\.168\.))(?!(172\.(1[6-9]|2[0-9]|3[0-1])\.))((([1-9]|1[1-9]|[2-9][0-9]|1[0-9]{2}|2[0-1][0-9]|22[0-3])\.)((0|1[0-9]{0,2}|[2-9][0-9]|2[0-4][0-9]|25[0-5])\.){2}(1[0-9]{0,2}|[2-9][0-9]|2[0-4][0-9]|25[0-4]))`, 0)
	if err != nil {
		fmt.Errorf("Error compiling IPv4 Regex: %s", err)
		return
	}
}

// Detect attempts to detect and parse an IPv4 address from the given input sample.
func Detect(sample string) (IPv4, error) {
	// 1. Assume IPv4 string is defanged. If defanged, it should be refanged.
	//
	// Original sample that is received by this function will be stored in 'sample'.
	// A working, refanged sample will be stored in `workingSample`.
	workingSample, _ := defang.Refang(sample)

	// 2. Read the entire string and check if it's an IPv4 (Any IP sequence in the string).
	//
	// Run IPv4 Regex check on the string.
	match, err := IPv4Regex.FindStringMatch(workingSample)

	// Return an error if there is a parse error.
	if err != nil {
		fmt.Errorf("Error parsing string: %s", workingSample)
		return IPv4{}, err
	}
	// If regex matched an IPv4.
	if match != nil {
		fanged := match.String()
		defanged, _ := defang.IPv4(fanged)
		DetectedIPv4 := IPv4{
			Fanged:   fanged,
			Defanged: defanged,
			Original: sample,
			Encoded:  url.QueryEscape(fanged),
		}

		fmt.Printf("\nFound match: %s, %s, %s, %s", DetectedIPv4.Fanged, DetectedIPv4.Defanged, DetectedIPv4.Original, DetectedIPv4.Encoded)
		// Return a parsed IPv4.
		return DetectedIPv4, nil
	}

	// If no matches are found, return an empty item and no error.
	return IPv4{}, nil
}
