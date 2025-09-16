package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
)

// ParseArgs converts a slice of command-line arguments into a model.ParsedArgs.
// 
// It recognizes the following flags:
//   -h, --help            -> CmdHelp
//   -v, --version         -> CmdVersion
//   -u, --use <path>      -> CmdSearchWithIndex, sets IndexPath to <path>
//   -s, --save [<path>]   -> CmdSave, sets SavePath to <path> or to the current working
//                           directory if no path is provided
//   -m, --match <mode>    -> sets MatchMode to one of: "exact" -> Exact,
//                           "contains" -> Contains, "regex" -> Regex (default is Contains)
//   -l, --list-indexes    -> CmdShowIndex
//
// Any unrecognized arguments are collected as unknown arguments; when no explicit
// command is set and there is at least one unknown argument, the function sets
// Command to CmdSearch and uses the first unknown argument as SearchArg. If the
// command is CmdSearchWithIndex and unknown arguments exist, the first unknown
// argument is used as SearchArg.
//
// Returns a fully populated model.ParsedArgs. Note: when -s/--save is provided
// without a following path, the function calls os.Getwd and will terminate the
// program via log.Fatal if getting the current directory fails.
func ParseArgs(args []string) model.ParsedArgs {
	parsed := model.ParsedArgs{
		Command:   model.CmdNone,
		MatchMode: model.Contains, //base value, else exact
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
				switch mode {
				case "exact":
					parsed.MatchMode = model.Exact
				case "contains":
					parsed.MatchMode = model.Contains
				case "regex":
					parsed.MatchMode = model.Regex
				}
				i++
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
