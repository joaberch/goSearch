package main

import (
	"github.com/joaberch/goSearch/cmd"
	"github.com/joaberch/goSearch/internal/model"
	"github.com/joaberch/goSearch/utils"
	"os"
)

func main() { //Future - Display line number
	args := os.Args[1:] //first is 'gosearch'
	if len(args) == 0 {
		cmd.ShowHelp()
		return
	}

	parsed := utils.ParseArgs(args)

	switch parsed.Command {
	case model.CmdHelp:
		cmd.ShowHelp()
	case model.CmdVersion:
		cmd.ShowVersion()
	case model.CmdSave:
		cmd.SaveIndex(parsed.SavePath)
	case model.CmdSearchWithIndex:
		cmd.SearchWithIndex(parsed.SearchArg, parsed.IndexPath, parsed.MatchMode)
	case model.CmdSearch:
		cmd.Search(parsed.SearchArg, parsed.MatchMode)
	case model.CmdShowIndex:
		cmd.ListIndexes()
	default:
		cmd.ShowHelp()
	}
}
