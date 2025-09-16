package model

// ParsedArgs carries normalized CLI arguments parsed from os.Args.
type ParsedArgs struct {
	Command   CommandType
	SearchArg string
	IndexPath string
	SavePath  string
	MatchMode MatchMode
}
