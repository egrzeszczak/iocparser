package url

import (
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/jakewarren/defang"
)

/*
^(?:(?:(?:https?|ftp):)?\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z0-9\u00a1-\uffff][a-z0-9\u00a1-\uffff_-]{0,62})?[a-z0-9\u00a1-\uffff]\.)+(?:[a-z\u00a1-\uffff]{2,}\.?))(?::\d{2,5})?(?:[/?#]\S*)?$
from: https://gist.github.com/dperini/729294
*/

type URL struct {
	Fanged   string
	Defanged string
	Original string
}

var FullURLRegex *regexp2.Regexp

func init() {
	var err error
	FullURLRegex, err = regexp2.Compile(`^(?:(?:(?:https?|ftp):)?\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z0-9\u00a1-\uffff][a-z0-9\u00a1-\uffff_-]{0,62})?[a-z0-9\u00a1-\uffff]\.)+(?:[a-z\u00a1-\uffff]{2,}\.?))(?::\d{2,5})?(?:[/?#]\S*)?$`, 0)
	// Throw error if there is a regex compilation error
	if err != nil {
		fmt.Errorf("Error compiling Full URL Regex: %s", err)
		return
	}
}

// ??? string -> URL struct
func Detect(sample string) (URL, error) {

	//
	// 1. Assume URL string is defanged. If defanged it should be refanged
	//

	// Original sample that is received by this function will be stored in 'sample'
	// Working, refanged sample will be stored in `workingSample`
	workingSample, _ := defang.Refang(sample)

	//
	// 2. Read entire string and check if it's an URL (Link from start to finish)
	//

	// Run Full URL Regex check on string
	match, err := FullURLRegex.FindStringMatch(workingSample)
	// Return an error if there is a parse error
	if err != nil {
		fmt.Errorf("Error parsing string: %s", workingSample)
		return URL{}, err
	}
	// If regex matched a URL
	if match != nil {
		fanged := match.String()
		defanged, _ := defang.URL(fanged)
		DetectedURL := URL{
			Fanged:   fanged,
			Defanged: defanged,
			Original: sample,
		}
		// Return an URL
		return DetectedURL, nil
	}

	// If no matches are found, return an empty item and no error
	return URL{}, nil
}
