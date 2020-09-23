package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/orsetii/defuse/cmd"
	"github.com/orsetii/defuse/cmd/parse"
	"gopkg.in/alecthomas/kingpin.v2"
)

// @TODO port flags to kong instead of kingpin
var (
	// Kingpin
	app = kingpin.New("defuse", "A CS:GO analyzer to evaluate performance in matches.")
	// Flags
	//OpenFileList = kingpin.Flag("demo-list", "Filepath of a list of demo files.").String()
	demoList    = app.Flag("demos", "path of demos to analyze").Short('d').Required().String()
	BUFSIZEFlag = app.Flag("buf", "Size of buffer (MiB) used when reading demo files.").Default("1").Short('b').Int()
	Verbose     = app.Flag("verbose", "Enable verbose mode").Short('v').Bool()

	// misc
	wg sync.WaitGroup
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	// Get Correct Buffer Size defaulting to 1 MiB.
	// @TODO Use in buffering opening BUFSIZE := *BUFSIZEFlag * 1048576 // 1MiB = 1048576 bytes.

	// @TODO add file reading functionality for OpenFileList
	// @TODO if FileList or OpenFileList not declared, exit.

	// Receive slice of strings to filepaths
	ValidPaths, err := cmd.ValidatePaths(demOpenFromArgs(*demoList))
	if err != nil {
		os.Exit(1)
	}

	// For each valid file, open the file and begin parsing for each.
	// @TODO Look to add concurrency around here?
	cmd.PrintInfo("Beginning File Parsing...\n")
	for _, v := range ValidPaths {
		wg.Add(1)
		go func(path string) {
			f, err := os.Open(path)
			defer f.Close()
			defer wg.Done()
			if err == nil {

				if *Verbose {
					cmd.PrintInfo(fmt.Sprintf("Parsing demo at: %s\n", path))
				}
				if pErr := parse.ParseDemo(f, *Verbose); pErr != nil {
					cmd.PrintWarn(fmt.Sprintf("Unable to parse %s.\nError: %s", path, pErr))
				}
			}
		}(v)
	}
	wg.Wait()
	fmt.Printf("\n")
	cmd.PrintInfo("Finished.\n")

}
func demOpenFromArgs(fileList string) []string {
	if fileList != "" {
		fileList = strings.Replace(fileList, " ", "", -1)
	}
	fmt.Println(fileList)
	if strings.Contains(fileList, ",") {
		return strings.Split(fileList, ",")
	}
	return []string{fileList}
}
