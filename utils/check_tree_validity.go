package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"strings"
)

// CheckTreeValidity reports whether a TreeElement should be indexed.
// For directories, it returns false if the directory name is present in model.InvalidFolder with a true value; otherwise it returns true.
// For files, it compares the file path case-insensitively against extensions in model.ExtensionsAllowed and returns false if the path ends with an extension that is marked not allowed; otherwise it returns true.
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
