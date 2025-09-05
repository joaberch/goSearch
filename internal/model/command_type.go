package model

type CommandType int

const (
	CmdNone CommandType = iota
	CmdSearch
	CmdHelp
	CmdVersion
	CmdSave
	CmdUse
)
