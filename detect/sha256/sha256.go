package sha256

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
)

type SHA256 struct {
	Value    string
	Encoded  string
	Original string
}

var SHA256Regex *regexp2.Regexp

func init() {
	var err error

	SHA256Regex, err = regexp2.Compile(`\b[a-fA-F0-9]{64}\b`, 0)
	if err != nil {
		fmt.Errorf("Error compiling SHA256 Regex: %s", err)
		return
	}
}

func Detect(sample string) (SHA256, error) {

	match, err := SHA256Regex.FindStringMatch(sample)

	if err != nil {
		fmt.Errorf("Error parsing string: %s", sample)
		return SHA256{}, err
	}

	if match != nil {
		value := match.String()
		DetectedSHA256 := SHA256{
			Value:    value,
			Original: sample,
			Encoded:  url.QueryEscape(value),
		}
		return DetectedSHA256, nil
	}

	return SHA256{}, nil
}
