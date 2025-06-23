package treebuilder

import (
	"log"
	"os"
)

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []TreeElement
}

func GetFileTree(selected []string) TreeElement {
	treeRoot := TreeElement{
		Name:     "root",
		Path:     "",
		IsDir:    true,
		Children: []TreeElement{},
	}

	for _, element := range selected {
		info, err := os.Stat(element)
		if err != nil {
			log.Fatal(err)
		}

		treeElement := TreeElement{
			Name:  info.Name(),
			Path:  element,
			IsDir: info.IsDir(),
		}
		if info.IsDir() {
			treeElement.Children = recursiveTree(element)
		}

		treeRoot.Children = append(treeRoot.Children, treeElement)
	}
	return treeRoot
}

func recursiveTree(path string) []TreeElement {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var children []TreeElement

	for _, entry := range entries {
		childPath := path + "\\" + entry.Name()

		child := TreeElement{
			Name:  entry.Name(),
			Path:  childPath,
			IsDir: entry.IsDir(),
		}

		if entry.IsDir() {
			child.Children = recursiveTree(childPath)
		}
		children = append(children, child)
	}
	return children
}
