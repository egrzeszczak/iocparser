package main

import (
	"fmt"

	"github.com/egrzeszczak/iocparser/detect"
	"github.com/egrzeszczak/iocparser/detect/ipv4"
	"github.com/egrzeszczak/iocparser/detect/md5"
	"github.com/egrzeszczak/iocparser/detect/sha1"
	"github.com/egrzeszczak/iocparser/detect/sha256"
	"github.com/egrzeszczak/iocparser/detect/url"
	"github.com/egrzeszczak/iocparser/input"
	"github.com/egrzeszczak/iocparser/reader"
)

func main() {

	// 1. Read input args
	filePath, _, _ := input.GetArgs()

	// // 2. Read file given in args
	fileLines, err := reader.Read(filePath)
	if err != nil {
		fmt.Errorf("Error while reading file: %s", err)
		return
	}

	var IoCs []detect.IoC

	for _, fileLine := range fileLines {

		DetectedURL, _ := url.Detect(fileLine)
		if (DetectedURL != url.URL{}) {
			IoCs = append(IoCs, detect.IoC{
				Value: DetectedURL,
			})
		}

		DetectedSHA256, _ := sha256.Detect(fileLine)
		if (DetectedSHA256 != sha256.SHA256{}) {
			IoCs = append(IoCs, detect.IoC{
				Value: DetectedSHA256,
			})
		}

		DetectedSHA1, _ := sha1.Detect(fileLine)
		if (DetectedSHA1 != sha1.SHA1{}) {
			IoCs = append(IoCs, detect.IoC{
				Value: DetectedSHA1,
			})
		}

		DetectedMD5, _ := md5.Detect(fileLine)
		if (DetectedMD5 != md5.MD5{}) {
			IoCs = append(IoCs, detect.IoC{
				Value: DetectedMD5,
			})
		}

		DetectedIPv4, _ := ipv4.Detect(fileLine)
		if (DetectedIPv4 != ipv4.IPv4{}) {
			IoCs = append(IoCs, detect.IoC{
				Value: DetectedIPv4,
			})
		}

	}

	for _, ioc := range IoCs {
		fmt.Printf("%T, %s\n\n", ioc.Value, ioc.Value)
	}

}
