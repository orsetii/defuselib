package main

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/orsetii/defuse/cmd"
	"github.com/orsetii/defuse/cmd/parse"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Flags
	FileList     = kingpin.Flag("demos", "List of Demo Files to parse").Short('d').Required().String()
	OpenFileList = kingpin.Flag("demo-list", "Filepath of a list of demo files.").String()
	BUFSIZE      = kingpin.Flag("buf", "Size of buffer used when reading demo files.").Short('b').Int()
	Verbose      = kingpin.Flag("verbose", "Enable verbose mode").Short('v').Bool()

	// misc
	wg sync.WaitGroup
)

func main() {

	kingpin.Parse()

	// @TODO add file reading functionality for OpenFileList
	// @TODO if FileList or OpenFileList not declared, exit.

	// Receive slice of strings to filepaths
	ValidPaths, err := cmd.ValidatePaths(demOpenFromArgs(*FileList))
	if err != nil {
		os.Exit(1)
	}

	// For each valid file, open the file and begin parsing for each.
	// @TODO Look to add concurrency around here?
	for _, v := range ValidPaths {
		f, err := os.Open(v)
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			if err == nil {
				if pErr := parse.ParseDemo(f, *Verbose); pErr != nil {
					log.Printf("Unable to parse %s.\nError: %s", path, pErr)
				}
			}
		}(v)
	}
	wg.Wait()

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
