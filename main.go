package main

import (
	"Go-LocalSearchEngine/indexer"
	"Go-LocalSearchEngine/streamer"
	"Go-LocalSearchEngine/treebuilder"
	"fmt"
	"github.com/ncruces/zenity"
	"log"
)

func main() {
	//Step 1 - Get Files
	//Step 1.1 - UI Folder selection
	selected := selectFolder()

	//Step 1.2 Filter file, we don't want to process exe file or dll or iso
	res := treebuilder.GetFileTree(selected)

	//Step 2 - Read file in streaming and normalize content
	//Step 2.1 - Stream and normalize file content
	var filesData []indexer.DataFile
	for _, file := range res.ValidFiles {
		toName := streamer.Stream(file)                                            //Shortcut are seen as file and not as folder, prevent stackOverflow error
		filesData = append(filesData, indexer.DataFile{Path: file, Token: toName}) //Step 2.2 - Store normalized content (temporarily)
	}

	//Step 3 - Create Inverted index
	invertedIndex := indexer.BuildInvertedIndex(filesData)
	fmt.Println(invertedIndex)
	//Step 3.1 - Create XML (Be wary of depth, prevent infinite loop because of symbolic link and display the number of file to index)

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
	fmt.Println("Emplacement racine : ", file)
	return file
}
