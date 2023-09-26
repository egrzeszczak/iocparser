package md5

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
)

type MD5 struct {
	Value    string
	Encoded  string
	Original string
}

var MD5Regex *regexp2.Regexp

func init() {
	var err error
	MD5Regex, err = regexp2.Compile(`\b[a-fA-F0-9]{32}\b`, 0)
	if err != nil {
		fmt.Errorf("Error compiling MD5 Regex: %s", err)
		return
	}
}
func Detect(sample string) (MD5, error) {
	match, err := MD5Regex.FindStringMatch(sample)
	if err != nil {
		fmt.Errorf("Error parsing string: %s", sample)
		return MD5{}, err
	}
	if match != nil {
		value := match.String()
		DetectedMD5 := MD5{
			Value:    value,
			Original: sample,
			Encoded:  url.QueryEscape(value),
		}
		return DetectedMD5, nil
	}
	return MD5{}, nil
}
