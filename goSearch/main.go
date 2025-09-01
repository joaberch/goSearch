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
	saveFlag := flag.Bool("s", false, "save search result")
	longSaveFlag := flag.Bool("save", false, "save search result")
	flag.Parse()

	args := flag.Args()

	if *versionFlag || *longVersionFlag {
		cmd.ShowVersion()
		return
	}

	if *helpFlag || *longHelpFlag {
		cmd.ShowHelp()
		return
	}

	if *saveFlag || *longSaveFlag {
		if len(args) == 0 { //gosearch -s
			path, _ := os.Getwd()
			cmd.SaveIndex(path)
			return
		} else { //gosearch -s path
			cmd.SaveIndex(args[0])
			return
		}
	}

	word := args[0]
	cmd.Search(word)
}
