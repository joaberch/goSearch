package main

import (
	"Go-LocalSearchEngine/compresser"
	"Go-LocalSearchEngine/indexer"
	"Go-LocalSearchEngine/program"
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Local Search in your file by selecting a Folder!")
	fmt.Println("1. Select Folder")
	fmt.Println("2. Use last folder selected (if an index already exists)")
	fmt.Print("Enter your choice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "exit":
		return
	case "1":
		program.Create()
	case "2":
		program.Use()
	}

	//Step 5.1 - Create UI
	index, err := indexer.LoadIndexFromXML("index.xml")
	if err != nil {
		log.Fatal(err)
	}

	uier.Loop(index, program.Create)
}
