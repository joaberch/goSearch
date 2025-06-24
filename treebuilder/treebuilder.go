package treebuilder

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []TreeElement
}

type TreeResult struct {
	Tree       TreeElement
	ValidFiles []string
}

func GetFileTree(selected []string) TreeResult {
	var validFiles []string

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
			children, files := recursiveTree(element)
			treeElement.Children = children
			validFiles = append(validFiles, files...) //TODO - understand
		} else if isValidFile(element) {
			validFiles = append(validFiles, element)
		}

		treeRoot.Children = append(treeRoot.Children, treeElement)
	}
	return TreeResult{Tree: treeRoot, ValidFiles: validFiles}
}

func recursiveTree(path string) ([]TreeElement, []string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var children []TreeElement
	var validFiles []string

	for _, entry := range entries {
		childPath := path + "\\" + entry.Name()

		child := TreeElement{
			Name:  entry.Name(),
			Path:  childPath,
			IsDir: entry.IsDir(),
		}

		if entry.IsDir() {
			subChildren, subFiles := recursiveTree(childPath)
			child.Children = subChildren
			validFiles = append(validFiles, subFiles...) //TODO - understand
		} else if isValidFile(childPath) {
			validFiles = append(validFiles, childPath)
		}
		children = append(children, child)
	}
	return children, validFiles
}

func isValidFile(path string) bool {
	e := strings.ToLower(filepath.Ext(path))                                          //e = extension
	return !(e == ".bin" || e == ".exe" || e == ".dll" || e == ".iso" || e == ".lnk") //TODO - add other and optimize
}
