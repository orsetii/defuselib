package cmd

import (
	"fmt"
	"os"
)

// Modular function which opens a single demo file, returns
func demOpen(path string) (f *os.File, err error) {
	f, err = os.Open(path)
	return f, nil
}

func validatePaths(files []string) {
	var errFiles []string
	for _, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			errFiles = append(errFiles, v)
		}
	}
	errFilelen := len(errFiles)
	if errFilelen > 0 {
		for _, v := range files {
			fmt.Printf("Error finding a file located at %s.\n", v)
		}
		fmt.Printf("Failed to Locate %d of your %d requested files.\n Would you still like to continue? (y/N)", errFilelen, len(files))
		var resp rune
		fmt.Scanf("%c", &resp)
		if resp != 'y' {
			os.Exit(1)
		}
	}
}
