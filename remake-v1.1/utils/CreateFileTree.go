package utils

import (
	"log"
	"os"
	"path/filepath"
	"search/internal/model"
)

func CreateFileTree(path string) model.TreeElement {
	//Get each children
	treeRoot := model.TreeElement{
		Name:     filepath.Base(path),
		Path:     path,
		IsDir:    true,
		Children: []*model.TreeElement{},
	}

	var entries, err = os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
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
			treeRoot.Children = append(treeRoot.Children, &model.TreeElement{
				Name:     info.Name(),
				Path:     fullPath,
				IsDir:    false,
				Children: nil,
			})
		}
	}
	return treeRoot
}
