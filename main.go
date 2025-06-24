package main

import (
	"Go-LocalSearchEngine/compresser"
	"Go-LocalSearchEngine/indexer"
	"Go-LocalSearchEngine/streamer"
	"Go-LocalSearchEngine/treebuilder"
	"Go-LocalSearchEngine/uier"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Delete basic XML file when the program is stopped, the file can be > 1Go so we only keep the .gz
	defer compresser.DeleteXML()

	//Step 1 - Get Files
	//Step 1.1 - UI Folder selection
	selected := uier.SelectFolder()

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
	//Step 3.1 - Create XML (Be wary of depth and display the number of file to index)
	err := indexer.SaveIndexAsXml(invertedIndex, "index.xml")
	if err != nil {
		log.Fatal("Erreur lors de la sauvegarde du fichier XML", err)
	}

	//Step 4 - Save Index
	//Step 4.1 - Compress XML to store
	compresser.CompressXMLFile() //To Save, currently not required TODO - UI enable to choose if we want to select folder or just research in a pre-existent one

	//Step 5 - Create UI to search word (check if match, not if fully equals)

	//Step 5.1 - Create UI
	index, err := indexer.LoadIndexFromXML("index.xml")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Entrez un mot à rechercher (ou 'exit' pour quitter) :")
		fmt.Print("> ")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		if query == "exit" {
			break
		}

		found := false
		for token, files := range index {
			if strings.Contains(token, query) {
				found = true
				fmt.Printf("'%s' trouvé dans le/les fichiers :\n", token)
				for _, file := range files {
					fmt.Println(" -", file)
				}
			}
		}

		if !found {
			fmt.Printf("Aucun fichier trouvé pour la recherche %s\n", query)
		}
	}

	//Step 5.3 -
}
