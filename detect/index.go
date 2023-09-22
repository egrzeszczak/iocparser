package detect

import (
	"regexp"
)

func Detect(lines []string) []IoC {
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
			domainMatch := domainRegex.FindStringSubmatch(line)
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
