package main

import (
	"flag"
	"goSearch/cmd"
	"os"
)

func main() {
	versionFlag := flag.Bool("v", false, "print version string")
	longVersionFlag := flag.Bool("version", false, "print version string")
	helpFlag := flag.Bool("h", false, "print usage string")
	longHelpFlag := flag.Bool("help", false, "print usage string")
	flag.Parse()

	args := os.Args[1:]

	if *versionFlag || *longVersionFlag {
		cmd.ShowVersion()
		return
	}

	if *helpFlag || *longHelpFlag || len(args) == 0 {
		cmd.ShowHelp()
		return
	}

	word := args[0]
	cmd.Search(word)
}
