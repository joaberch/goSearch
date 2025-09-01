package main

import (
	"flag"
	"goSearch/cmd"
	"os"
)

func main() {
	//TODO - Use other flag parser
	versionFlag := flag.Bool("v", false, "print version string")
	longVersionFlag := flag.Bool("version", false, "print version string")
	helpFlag := flag.Bool("h", false, "print usage string")
	longHelpFlag := flag.Bool("help", false, "print usage string")
	saveFlag := flag.Bool("s", false, "save search result")
	longSaveFlag := flag.Bool("save", false, "save search result")
	flag.Parse()

	args := flag.Args()

	switch {
	case *versionFlag || *longVersionFlag:
		cmd.ShowVersion()
	case *helpFlag || *longHelpFlag:
		cmd.ShowHelp()
	case *saveFlag || *longSaveFlag:
		var path string
		if len(args) == 0 {
			path, _ = os.Getwd()
		} else {
			path = args[0]
		}
		cmd.SaveIndex(path)
	default:
		fmt.Println("default")
		cmd.Search(args)
	}
}
