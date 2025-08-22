package cmd

import "errors"

var ErrUnknownCommand = errors.New("unknown command")

// CommandsMap is a mapping of command names to their corresponding command types, defining available commands for lookup.
var CommandsMap = map[string]Type{
	"help":    CmdShowHelp,
	"version": CmdShowVersion,
}

// GetCommandType GetCommand retrieves the command type for a given name from CommandsMap, or returns an error if the command is unknown.
func GetCommandType(name string) (Type, error) {
	if cmdType, found := CommandsMap[name]; found {
		return cmdType, nil
	}
	return CmdSearch, nil
}
