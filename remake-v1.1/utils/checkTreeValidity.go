package utils

import (
	"search/internal/model"
	"strings"
)

func CheckTreeValidity(entry model.TreeElement) bool {
	if entry.IsDir {
		//TODO : choose which directory not process
	} else {
		for ext, isValid := range model.InvalidExtensions {
			if strings.HasSuffix(entry.Path, ext) && !isValid {
				return false
			}
		}
	}
	return true
}
