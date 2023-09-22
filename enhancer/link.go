package enhancer

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

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
	// Read and parse the cti.conf file.
	configFile, err := os.Open("cti.conf")
	if err != nil {
		fmt.Println("Error opening cti.conf:", err)
		return
	}
	defer configFile.Close()

	// Unmarshal the JSON into the Config struct
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
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
