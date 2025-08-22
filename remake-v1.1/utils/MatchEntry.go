package utils

import (
	"encoding/xml"
	"strings"
)

func MatchEntry(entry xml.StartElement, word string) bool {
	for _, attr := range entry.Attr { //foreach attribute
		if strings.Contains(attr.Value, word) { //if the attribute contains the word wanted
			return true
		}
	}
	return false
}
