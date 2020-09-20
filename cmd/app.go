package cmd

import (
	"strings"

	"github.com/fatih/color"
)

func PrintWarn(content string) {
	red := color.New(color.Bold, color.FgRed).PrintfFunc()
	red("\n%s", content)
}

func PrintInfo(content string) {
	info := color.New(color.FgCyan).PrintfFunc()
	info("\n[*] %s", content)
}

func PrintQuestion(content string) {
	info := color.New(color.FgHiYellow).PrintfFunc()
	info("\n%s", content)
}
func PrintProg(progInt int) {
	prog := color.New(color.FgHiWhite).PrintfFunc()
	prog("%s>\r", strings.Repeat("=", progInt))
}
