package main

import (
	"fmt"
	"os"
	"search/cmd"
)

func main() {
	userCommand, err := cmd.GetCommandType(os.Args[1])
	if err != nil {
		fmt.Println(err)
		cmd.ShowHelp()
		return
	}

	switch userCommand {
	case cmd.CmdShowVersion:
		cmd.ShowVersion()
		break
	case cmd.CmdSearch:
		word := os.Args[1]
		cmd.Search(word)
		break
	default:
		cmd.ShowHelp()
	}
}
