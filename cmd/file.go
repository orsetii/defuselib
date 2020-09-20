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

			PrintWarn(fmt.Sprintf("\n\tError finding a file located at %s.\n", v))
			errdFiles++
		} else {
			FoundFiles = append(FoundFiles, v)
		}
	}

	FoundFileNum := len(FoundFiles)

	switch FoundFileNum {
	case len(files):
		PrintInfo(fmt.Sprintf("All %d files validated.\n", FoundFileNum))
		return FoundFiles, nil
	case 0:
		PrintInfo(fmt.Sprintf("Could not validate any files."))
		PrintInfo(fmt.Sprintf("Exiting...\n\n"))
		os.Exit(1)
	default:
		PrintQuestion(fmt.Sprintf("\nFailed to Locate %d of your %d requested files.\n Would you still like to continue? (y/N)", errdFiles, len(files)))
	}
	var resp string
	fmt.Scanf("%s\n", &resp)
	if resp == "y" {

		return FoundFiles, nil
	}
	return nil, errors.New("User Requested Exit")
}
