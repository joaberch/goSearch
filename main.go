package main

import (
	"Go-LocalSearchEngine/treebuilder"
	"fmt"
	"github.com/ncruces/zenity"
	"log"
)

func main() {
	//Step 1 - Get Files
	//Step 1.1 - UI Folder selection
	selected := selectFolder()
	tree := treebuilder.GetFileTree(selected)
	fmt.Println(tree)
	printTree(tree, 0)

	//Step 2 - Read file in streaming and normalize content
	//Step 2.1 - Filter (exe, etc)

	//Step 2.2 - Stream file content

	//Step 2.3 - Normalize text content

	//Step 2.4 - Store normalized content (temporarily)

	//Step 3 - Create Inverted index
	//Step 3.1 - Create XML (Be wary of depth, filter extension like exe to save resources, prevent infinite loop because of symbolic link and display the number of file to index)

	//Step 4 - Save Index
	//Step 4.1 - Compress XML to store
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
