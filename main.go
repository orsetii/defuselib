package main

import (
	"flag"
	"strings"

	"github.com/orsetii/counteranalysis/cmd"
)

var (
	fileList string
)

func main() {
	fileList := flag.String("demo", "", "Demo file paths, split by commas.") // Make this a custom flag, with a call to demOpenFromArgs in the Parsing functionality
	flag.Parse()
	if *fileList != "" {
		cmd.Start(demOpenFromArgs(*fileList))
	}
}

func demOpenFromArgs(fileList string) []string {
	if strings.Contains(fileList, ",") {
		return strings.Split(fileList, ",")
	}
	return []string{fileList}
}
