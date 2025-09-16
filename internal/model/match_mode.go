package model

type MatchMode int

const (
	// Exact matches whole key
	Exact MatchMode = iota
	// Contains does substring match
	Contains
	// Regex matches regex expression
	Regex
)
