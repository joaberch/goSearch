package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
	"path/filepath"
)

// CreateFileTree builds a recursive TreeElement structure from the given directory path.
func CreateFileTree(path string) model.TreeElement {
	treeRoot := model.TreeElement{
		Name:     filepath.Base(path),
		Path:     path,
		IsDir:    true,
		Children: []*model.TreeElement{},
	}

	var entries, err = os.ReadDir(path)
	if err != nil {
		log.Printf("Error reading directory %s: %s", path, err)
		return treeRoot
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		info, err := entry.Info()
		if err != nil {
			log.Printf("Error reading directory %s: %s", path, err)
			continue
		}

		element := model.TreeElement{
			Name:  info.Name(),
			Path:  fullPath,
			IsDir: entry.IsDir(),
		}

		if !CheckTreeValidity(element) {
			continue
		}

		if entry.IsDir() {
			child := CreateFileTree(fullPath)
			treeRoot.Children = append(treeRoot.Children, &child)
		} else {
			treeRoot.Children = append(treeRoot.Children, &element)
		}
	}
	return treeRoot
}
