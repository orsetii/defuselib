package main

import (
	"flag"
	"strings"

	"github.com/orsetii/defuse/cmd"
)

var (
	fileList string
)

func main() {
	fileListFlag := flag.String("demo", "", "Demo file paths, split by commas.") // Make this a custom flag, with a call to demOpenFromArgs in the Parsing functionality
	flag.Parse()
	cmd.Start(demOpenFromArgs(*fileListFlag))
}

func demOpenFromArgs(fileList string) []string {
	if fileList != "" {
		fileList = strings.Replace(fileList, " ", "", -1)
	}
	if strings.Contains(fileList, ",") {

		return strings.Split(fileList, ",")
	}
	return []string{fileList}
}
