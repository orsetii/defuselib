package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Main(filePaths []string) {
	// Receive slice of strings to filepaths
	ValidPaths, err := validatePaths(filePaths)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(ValidPaths) //TODO this is here for preventing var not used err
}

func printWarn(content string) {
	red := color.New(color.Bold, color.FgRed).PrintfFunc()
	red("%s", content)
}

func printInfo(content string) {
	info := color.New(color.FgCyan).PrintfFunc()
	info("[*] %s", content)
}

func printQuestion(content string) {
	info := color.New(color.FgHiYellow).PrintfFunc()
	info("%s", content)
}
