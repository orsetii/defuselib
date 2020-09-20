package cmd

import (
	"github.com/fatih/color"
)

func PrintWarn(content string) {
	red := color.New(color.Bold, color.FgRed).PrintfFunc()
	red("%s", content)
}

func PrintInfo(content string) {
	info := color.New(color.FgCyan).PrintfFunc()
	info("[*] %s", content)
}

func PrintQuestion(content string) {
	info := color.New(color.FgHiYellow).PrintfFunc()
	info("%s", content)
}
