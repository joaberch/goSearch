package model

type CommandType int

const (
	// CmdNone indicates no command parsed.
	CmdNone CommandType = iota
	// CmdSearch runs an on-the-fly search (no index).
	CmdSearch
	// CmdHelp prints usage/help.
	CmdHelp
	// CmdVersion prints version info.
	CmdVersion
	// CmdSave builds and save the index.
	CmdSave
	// CmdSearchWithIndex searches using a saved index.
	CmdSearchWithIndex
	// CmdShowIndex shows the list of indexes available.
	CmdShowIndex
)
