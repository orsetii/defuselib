package cmd

import (
	"errors"
	"fmt"
	"os"
)

func ValidatePaths(files []string) (FoundFiles []string, err error) {
	PrintInfo("Validating Files...\n")
	errdFiles := 0
	for _, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {

			PrintWarn(fmt.Sprintf("\n\tError finding a file located at %s.", v))
			errdFiles++
		} else {
			FoundFiles = append(FoundFiles, v)
		}
	}
	if len(files) == len(FoundFiles) {
		return FoundFiles, nil
	}
	PrintQuestion(fmt.Sprintf("\n\nFailed to Locate %d of your %d requested files.\n Would you still like to continue? (y/N)", errdFiles, len(files)))
	var resp string
	fmt.Scanf("%s\n", &resp)
	if resp == "y" {

		return FoundFiles, nil
	}
	return nil, errors.New("User Requested Exit")
}
