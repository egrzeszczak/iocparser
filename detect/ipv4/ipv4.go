package ipv4

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
	"github.com/jakewarren/defang"
)

type IPv4 struct {
	Fanged   string
	Defanged string
	Encoded  string
	Original string
}

var IPv4Regex *regexp2.Regexp

func init() {
	var err error
	IPv4Regex, err = regexp2.Compile(`(\d+)(?<!10)(?<!127)\.(\d+)(?<!192\.168)(?<!172\.(1[6-9]|2[0-9]|3[0-1]))\.(\d+)\.(\d+)`, 0)
	if err != nil {
		fmt.Errorf("Error compiling IPv4 Regex: %s", err)
		return
	}
}
func Detect(sample string) (IPv4, error) {
	workingSample, _ := defang.Refang(sample)

	match, err := IPv4Regex.FindStringMatch(workingSample)

	if err != nil {
		fmt.Errorf("Error parsing string: %s", workingSample)
		return IPv4{}, err
	}

	if match != nil {
		fanged := match.String()
		defanged, _ := defang.IPv4(fanged)
		DetectedIPv4 := IPv4{
			Fanged:   fanged,
			Defanged: defanged,
			Original: sample,
			Encoded:  url.QueryEscape(fanged),
		}
		return DetectedIPv4, nil
	}

	return IPv4{}, nil
}
