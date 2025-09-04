package main

import (
	"github.com/joaberch/goSearch/cmd"
	"os"
)

func main() { //Future - Display line number
	args := os.Args[1:] //1 is gosearch
	if len(args) == 0 {
		cmd.ShowHelp()
		return
	} else if len(args) == 1 {
		cmd.Search(args[0])
		return
	}

	for i, arg := range args {
		switch arg {
		case "-h", "--help":
			cmd.ShowHelp()
			return
		case "-v", "--version":
			cmd.ShowVersion()
			return
		case "-u", "--use":
			if i+2 < len(args) {
				cmd.SearchWithIndex(args[i+2], args[i+1])
				return
			}
		case "-s", "--save":
			if i+1 < len(args) {
				cmd.SaveIndex(args[i+1])
				return
			}
		default:
			cmd.ShowHelp()
			return
		}
	}
}
