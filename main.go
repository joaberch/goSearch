package main

import (
	"Go-LocalSearchEngine/treebuilder"
	"bufio"
	"fmt"
	"github.com/ncruces/zenity"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	//Step 1 - Get Files
	//Step 1.1 - UI Folder selection
	selected := selectFolder()

	//Step 1.2 Filter file, we don't want to process exe file or dll or iso
	res := treebuilder.GetFileTree(selected)

	//Step 2 - Read file in streaming and normalize content
	//Step 2.1 - Stream and normalize file content

	//Normalization guideline :
	//1 - Everything in lowercase
	//2 - Remove punctuation
	//3 - Remove special character
	//4 - Remove stop words (the, is, and, etc) - not required but surely needed in that application
	//5 - Stemming
	//6 - Lemmatization

	for _, file := range res.ValidFiles {
		toName := stream(file)
		fmt.Println(toName)
	}

	//Step 2.2 - Store normalized content (temporarily)

	//Step 3 - Create Inverted index
	//Step 3.1 - Create XML (Be wary of depth, filter extension like exe to save resources, prevent infinite loop because of symbolic link and display the number of file to index)

	//Step 4 - Save Index
	//Step 4.1 - Compress XML to store
}

func stream(path string) []string {
	//Stream
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //Line length CAN be longer than 65536, TODO - see if the Buffer Method is required
	var tokens []string
	for scanner.Scan() {
		line := scanner.Text()
		normalizedLine := normalize(line)
		tokens = append(tokens, normalizedLine...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}

func normalize(line string) []string {
	//1 - Lowercase
	line = strings.ToLower(line)

	//2 & 3 - Punctuation & Special character
	regxp, _ := regexp.Compile(`[^\w\s]`)
	line = regxp.ReplaceAllString(line, "")

	//Split for each character
	words := strings.Fields(line)

	stopWords := map[string]bool{
		"the": true, "is": true, "and": true, "a": true, "to": true, "in": true, "of": true, "their": true, "theirs": true, "de": true,
	}

	//4 - Remove stop words
	var filtered []string
	for _, word := range words {
		if !stopWords[word] {
			filtered = append(filtered, word)
		}
	}

	//5 - Stemming and Lemmatization

	return filtered
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
