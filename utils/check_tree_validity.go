package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"strings"
)

// CheckTreeValidity checks if the TreeElement is valid for indexing
func CheckTreeValidity(entry model.TreeElement) bool {
	if entry.IsDir {
		if invalid, ok := model.InvalidFolder[entry.Name]; ok && invalid {
			return false
		}
		return true
	} else {
		lowerPath := strings.ToLower(entry.Path)
		for ext, allowed := range model.ExtensionsAllowed {
			if strings.HasSuffix(lowerPath, strings.ToLower(ext)) && !allowed {
				return false
			}
		}
		return true
	}
}
