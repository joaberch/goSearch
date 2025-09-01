package model

// InvalidWord maps word with their invalidity
var InvalidWord = map[string]bool{
	"<UNK>":  true,
	"the":    true,
	"is":     true,
	"and":    true,
	"their":  true,
	"theirs": true,
	"a":      true,
	"to":     true,
	"in":     true,
	"of":     true,
	"de":     true,
}
