package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
)

func ParseArgs(args []string) model.ParsedArgs {
	parsed := model.ParsedArgs{
		Command:   model.CmdNone,
		MatchMode: "contains", //base value, else exact
	}

	var unknownArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			parsed.Command = model.CmdHelp
		case "-v", "--version":
			parsed.Command = model.CmdVersion
		case "-u", "--use":
			parsed.Command = model.CmdSearchWithIndex
			if i+1 < len(args) {
				parsed.IndexPath = args[i+1]
				i++
			}
		case "-s", "--save":
			parsed.Command = model.CmdSave
			if i+1 < len(args) {
				parsed.SavePath = args[i+1]
				i++
			} else {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				parsed.SavePath = path
			}
		case "-m", "--match":
			if i+1 < len(args) {
				mode := args[i+1]
				if mode == "exact" || mode == "contains" {
					parsed.MatchMode = mode
					i++
				} else {
					parsed.Command = model.CmdHelp
				}
			}
		case "-l", "--list-indexes":
			parsed.Command = model.CmdShowIndex
		default:
			unknownArgs = append(unknownArgs, arg) //Get search word, different position possible
		}
	}

	//Search
	if parsed.Command == model.CmdNone && len(unknownArgs) > 0 {
		parsed.Command = model.CmdSearch
		parsed.SearchArg = unknownArgs[0]
	} else if parsed.Command == model.CmdSearchWithIndex && len(unknownArgs) > 0 {
		parsed.SearchArg = unknownArgs[0]
	}

	return parsed
}
