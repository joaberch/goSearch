package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"strings"
)

// CheckTreeValidity checks if the TreeElement is valid for indexing
func CheckTreeValidity(entry model.TreeElement) bool {
	if entry.IsDir {
		for folder, isInvalid := range model.InvalidFolder {
			if entry.Name == folder && isInvalid {
				return false
			}
		}
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
