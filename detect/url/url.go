package url

import (
	"fmt"
	"net/url"

	"github.com/dlclark/regexp2"
	"github.com/jakewarren/defang"
)

type URL struct {
	Fanged   string
	Defanged string
	Original string
	Encoded  string
}

var URLRegex *regexp2.Regexp

func init() {
	var err error
	URLRegex, err = regexp2.Compile(`^(?:(?:(?:https?|ftps?):)?\/\/)?(?:[^\r\n\t\f\v@ ]+(?::[^\r\n\t\f\v@ ]*)?)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z0-9\u00a1-\uffff][a-z0-9\u00a1-\uffff-]{0,62})?[a-z0-9\u00a1-\uffff]\.)+(?:[a-z\u00a1-\uffff]{2,}\.?))(?::\d{2,5})?(?:[/?#][^\r\n\t\f\v@ ]*)?$`, 0)
	if err != nil {
		fmt.Errorf("Error compiling Full URL Regex: %s", err)
		return
	}
}
func Detect(sample string) (URL, error) {
	workingSample, _ := defang.Refang(sample)

	match, err := URLRegex.FindStringMatch(workingSample)

	if err != nil {
		fmt.Errorf("Error parsing string: %s", workingSample)
		return URL{}, err
	}
	if match != nil {
		fanged := match.String()
		defanged, _ := defang.URL(fanged)
		DetectedURL := URL{
			Fanged:   fanged,
			Defanged: defanged,
			Original: sample,
			Encoded:  url.QueryEscape(fanged),
		}

		return DetectedURL, nil
	}

	return URL{}, nil
}
