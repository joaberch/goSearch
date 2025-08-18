package cmd

type Type int

const (
	CmdShowHelp Type = iota
	CmdShowVersion
	CmdSearch
	CmdDisplayTree
	CmdIndexate
)
