package uier

import (
	"Go-LocalSearchEngine/indexer"
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
		query = strings.ToLower(query)

		if query == "!exit" {
			break
		} else if query == "!select" {
			onSelect()
			newIndex, err := indexer.LoadIndexFromXML("index.xml")
			if err != nil {
				return
			}
			index = newIndex
		} else if query == "!displayTree" {
			printTree(index)
		} else if query == "!help" {
			fmt.Println("\nCommandes disponibles :")
			fmt.Println(" !help        → Affiche ce message d’aide")
			fmt.Println(" !exit        → Quitte le programme")
			fmt.Println(" !select      → Sélectionne un ou plusieurs dossiers à indexer pour effectuer la recherche dedans")
			fmt.Println(" !displayTree → WIP - Affiche l'arborescence")
			fmt.Println(" <mot>        → Recherche un mot dans les fichiers indexés\n")
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
func printTree(index map[string][]string) {
	log.Println("Cette fonction n'affiche pas forcément les fichiers dans leur dossier respectif")
	seenFiles := make(map[string]bool)
	seenDirs := make(map[string]bool)

	for _, files := range index {
		for _, file := range files {
			if seenFiles[file] {
				continue
			}
			seenFiles[file] = true

			parts := strings.Split(file, string(os.PathSeparator))
			var pathBuilder []string

			for i := 0; i < len(parts); i++ {
				pathBuilder = append(pathBuilder, parts[i])
				currentPath := strings.Join(pathBuilder, string(os.PathSeparator))

				if i == len(parts)-1 {
					//File
					fmt.Print(strings.Repeat(" ", i))
					fmt.Printf("📄 %s\n", parts[i])
				} else {
					if seenDirs[currentPath] {
						continue
					}
					seenDirs[currentPath] = true

					fmt.Print(strings.Repeat(" ", i))
					fmt.Printf("📁 %s\n", parts[i])
				}
			}
		}
	}
	log.Println("Cette fonction n'affiche pas forcément les fichiers dans leur dossier respectif")
}

// SelectFolder Select one or multiple folder
func SelectFolder() []string {
	file, err := zenity.SelectFileMultiple(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}
	return file
}
