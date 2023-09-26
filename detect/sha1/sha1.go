package sha1

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
)

type SHA1 struct {
	Value    string
	Encoded  string
	Original string
}

var SHA1Regex *regexp2.Regexp

func init() {
	var err error

	SHA1Regex, err = regexp2.Compile(`\b[a-fA-F0-9]{40}\b`, 0)
	if err != nil {
		fmt.Errorf("Error compiling SHA1 Regex: %s", err)
		return
	}
}

func Detect(sample string) (SHA1, error) {

	match, err := SHA1Regex.FindStringMatch(sample)

	if err != nil {
		fmt.Errorf("Error parsing string: %s", sample)
		return SHA1{}, err
	}

	if match != nil {
		value := match.String()
		DetectedSHA1 := SHA1{
			Value:    value,
			Original: sample,
			Encoded:  url.QueryEscape(value),
		}
		return DetectedSHA1, nil
	}

	return SHA1{}, nil
}
