package model

type ParsedArgs struct {
	Command   CommandType
	SearchArg string
	IndexPath string
	SavePath  string
	MatchMode string
}
