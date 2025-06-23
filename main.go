package main

import (
	"Go-LocalSearchEngine/treebuilder"
	"fmt"
	"github.com/ncruces/zenity"
	"log"
)

func main() {
	selected := selectFolder()
	tree := treebuilder.GetFileTree(selected)
	fmt.Println(tree)
	printTree(tree, 0)
}

// Debug func
func printTree(tree treebuilder.TreeElement, depth int) {
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

// Select one or multiple folder
func selectFolder() []string {
	file, err := zenity.SelectFileMultiple(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Emplacement : ", file)
	return file
}
