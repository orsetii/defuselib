package cmd

import (
	"fmt"
)

func Start(filePaths []string) {
	// Receive slice of strings to filepaths
	fmt.Printf("%#v\n", filePaths)
	validatePaths(filePaths)
}
