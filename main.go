package main

import (
	"strings"

	"github.com/orsetii/defuse/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Flags
	FileList     = kingpin.Flag("demos", "List of Demo Files to parse").Short('d').Required().String()
	OpenFileList = kingpin.Flag("demo-list", "Filepath of a list of demo files.").String()
	BUFSIZE      = kingpin.Flag("buf", "Size of buffer used when reading demo files.").Short('b').Int()
	Verbose      = kingpin.Flag("verbose", "Enable verbose mode").Short('v').Bool()
)

func main() {
	// TODO add file reading functionality for OpenFileList
	// TODO if FileList or OpenFileList not declared, exit.
	kingpin.Parse()
	cmd.Main(demOpenFromArgs(*FileList))

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
