package utils

import (
	"github.com/joaberch/Go-LocalSearchEngine/internal/model"
	"strings"
)

// CheckTreeValidity checks if the TreeElement is valid for indexing
func CheckTreeValidity(entry model.TreeElement) bool {
	if entry.IsDir {
		//FUTURE : choose which directory not process, like node_modules or .git
		return true
	} else {
		for ext, isValid := range model.InvalidExtensions {
			if strings.HasSuffix(entry.Path, ext) && !isValid {
				return false
			}
		}
		return true
	}
}
