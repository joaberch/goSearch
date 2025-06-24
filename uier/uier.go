package uier

import (
	"Go-LocalSearchEngine/indexer"
	"Go-LocalSearchEngine/treebuilder"
	"bufio"
	"fmt"
	"github.com/ncruces/zenity"
	"log"
	"os"
	"strings"
)

func Loop(index map[string][]string, onSelect func()) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Mot à rechercher (!help, !exit ou !select) :")
		fmt.Print("> ")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		if query == "!exit" {
			break
		} else if query == "!select" {
			onSelect()
			newIndex, err := indexer.LoadIndexFromXML("index.xml")
			if err != nil {
				return
			}
			index = newIndex
		} else if query == "!help" {

			fmt.Println("\nCommandes disponibles :")
			fmt.Println(" !help   → Affiche ce message d’aide")
			fmt.Println(" !exit   → Quitte le programme")
			fmt.Println(" !select → Sélectionne un ou plusieurs dossiers à indexer pour effectuer la recherche dedans")
			fmt.Println(" <mot>   → Recherche un mot dans les fichiers indexés\n")

		} else {
			found := false
			for token, files := range index {
				if strings.Contains(token, query) {
					found = true
					fmt.Printf("%s trouvé dans le/les fichiers :\n", token)
					for _, file := range files {
						fmt.Println(" -", file)
					}
				}
			}

			if !found {
				fmt.Printf("Aucun fichier trouvé pour la recherche %s\n", query)
			}
		}
	}
}

// Debug func
func PrintTree(tree treebuilder.TreeElement, depth int) {
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
		PrintTree(child, depth+1)
	}
}

// Select one or multiple folder
func SelectFolder() []string {
	file, err := zenity.SelectFileMultiple(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}
	return file
}
