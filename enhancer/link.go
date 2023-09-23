package enhancer

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/egrzeszczak/iocparser/detect"
	"github.com/egrzeszczak/iocparser/help"
)

type Link struct {
	Name string
	URL  string
}

type Config struct {
	Services map[string]struct {
		Links map[string][]string `json:"Links"`
	} `json:"Services"`
}

var config Config

func init() {
	configContent := `
	{
		"Services": {
			"VirusTotal": {
				"Links": {
					"https://www.virustotal.com/gui/search/": [
						"MD5", "SHA1", "SHA256"
					],
					"https://www.virustotal.com/gui/ip-address/": [
						"IPv4"
					],
					"https://www.virustotal.com/gui/domain/": [
						"Domain"
					]
				}
			},
			"Talos": {
				"Links": {
					"https://www.talosintelligence.com/reputation_center/lookup?search=": [
						"Domain", "IPv4", "Email"
					]
				}
			},
			"XForce": {
				"Links": {
					"https://exchange.xforce.ibmcloud.com/url/": [
						"URL"
					],
					"https://exchange.xforce.ibmcloud.com/ip/": [
						"IPv4"
					],
					"https://exchange.xforce.ibmcloud.com/malware/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"URLHaus": {
				"Links": {
					"https://urlhaus.abuse.ch/browse.php?search=": [
						"URL"
					]
				}
			},
			"AbuseIPDB": {
				"Links": {
					"https://www.abuseipdb.com/check/": [
						"IPv4", "Domain"
					]
				}
			},
			"InQuest": {
				"Links": {
					"https://labs.inquest.net/search/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"Bazaar": {
				"Links": {
					"https://bazaar.abuse.ch/sample/": [
						"SHA256"
					]
				}
			},
			"MWDB CERT": {
				"Links": {
					"https://mwdb.cert.pl/sample/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"Hybrid Analysis": {
				"Links": {
					"https://www.hybrid-analysis.com/sample/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"Malprob": {
				"Links": {
					"https://malprob.io/report/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"AnyRun": {
				"Links": {
					"https://app.any.run/submissions/#filehash:": [
						"MD5", "SHA1", "SHA256"
					]
				}
			},
			"Yomi": {
				"Links": {
					"https://yomi.yoroi.company/submissions/": [
						"MD5", "SHA1", "SHA256"
					]
				}
			}
		}
	}
	`

	// Unmarshal the JSON content into the Config struct
	if err := json.Unmarshal([]byte(configContent), &config); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetConfig() Config {
	return config
}

func CreateLinks(ioc detect.IoC) []Link {
	var links []Link

	// Iterate through the "Services" map
	for serviceName, service := range config.Services {
		// Iterate through the "Links" map for each service
		for linkName, values := range service.Links {
			if help.StringInSlice(ioc.Type, values) {
				links = append(links, Link{
					Name: serviceName,
					URL:  linkName + url.QueryEscape(ioc.Value),
				})
			}
			// // Iterate through the slice of values for each link
			// for _, value := range values {
			// 	fmt.Println("    Value:", value)
			// }
		}
	}

	return links
}
