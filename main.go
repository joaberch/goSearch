package main

import (
	"fmt"
	"github.com/ncruces/zenity"
	"log"
	"os"
)

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []TreeElement
}

func main() {
	selected := selectF()
	tree := getFileTree(selected)
	printTree(tree, 0)
}

func printTree(tree TreeElement, depth int) {
	indentation := ""
	for i := 0; i < depth; i++ {
		indentation += " "
	}

	if tree.IsDir {
		fmt.Printf("%s📁 %s\n", indentation, tree.Name)
	} else {
		fmt.Printf("%s📄 %s\n", indentation, tree.Name)
	}

	for _, child := range tree.Children {
		printTree(child, depth+1)
	}
}

func selectF() []string {
	file, err := zenity.SelectFileMultiple(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Emplacement : ", file)
	return file
}

func getFileTree(selected []string) TreeElement {
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
