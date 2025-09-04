package main

import (
	"github.com/joaberch/goSearch/cmd"
	"log"
	"os"
)

func main() { //Future - Display line number
	args := os.Args[1:] //first is gosearch
	if len(args) == 0 {
		cmd.ShowHelp()
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
			var path string
			var err error
			if i+1 < len(args) {
				path = args[i+1]
			} else {
				path, err = os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
			}
			cmd.SaveIndex(path)
			return
		default:
			if len(args) == 1 {
				cmd.Search(args[0])
				return
			}
			cmd.ShowHelp()
			return
		}
	}
}
