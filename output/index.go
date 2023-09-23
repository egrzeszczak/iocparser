package output

import (
	"fmt"
	"sort"

	"github.com/egrzeszczak/iocparser/detect"
	"github.com/egrzeszczak/iocparser/enrich/reputation"
)

// Print IoCs to Markdown style function
func OutputMarkdown(iocs []detect.IoC) {
	// Create a map to store grouped and unique IoCs by Type.
	groupedUniqueIoCs := make(map[string]map[string]struct{})
	var typeOrder []string

	for _, ioc := range iocs {
		// Create a map for each Type if it doesn't exist.
		if _, exists := groupedUniqueIoCs[ioc.Type]; !exists {
			groupedUniqueIoCs[ioc.Type] = make(map[string]struct{})
			typeOrder = append(typeOrder, ioc.Type) // Add the detected Type to typeOrder.
		}

		// Add the IoC Value to the map to ensure uniqueness.
		groupedUniqueIoCs[ioc.Type][ioc.Value] = struct{}{}
	}

	// Sort the typeOrder slice alphabetically.
	sort.Strings(typeOrder)

	// Print the grouped, sorted, and unique IoCs.
	for _, t := range typeOrder {
		if ioCsOfType, exists := groupedUniqueIoCs[t]; exists {
			// Collect the IoCs for this type.
			var sortedIoCs []string
			for ioc := range ioCsOfType {
				sortedIoCs = append(sortedIoCs, ioc)
			}

			// Sort the IoCs alphabetically.
			sort.Strings(sortedIoCs)

			// Print the IoCs.
			fmt.Printf("## %s\n\n", t)
			for _, ioc := range sortedIoCs {
				fmt.Printf("- `%s`\n", ioc)
				// Create links for each IoC value.
				links := reputation.Create(detect.IoC{Value: ioc, Type: t})

				fmt.Printf("\n\t - Check on: ")
				for _, link := range links {
					fmt.Printf("[%s](%s), ", link.Name, link.URL)
				}
				fmt.Println()
			}
			fmt.Println() // Add a blank line between groups.
		}
	}
}
