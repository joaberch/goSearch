package utils

import (
	"bufio"
	"github.com/joaberch/Go-LocalSearchEngine/internal/model"
	"log"
	"os"
	"strings"
)

// StreamFile reads a file line by line, normalizes its content, and returns a FileData structure containing the extracted words.
func StreamFile(element *model.TreeElement) model.FileData {
	newFile := model.FileData{
		Path: element.Path,
	}

	file, err := os.Open(newFile.Path)
	if err != nil {
		log.Printf("Error opening file %s : %v", newFile.Path, err)
		return newFile
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)    // Error if the line length is longer than 65'536 character so we define a max capacity
	const maxCapacity = 1024 * 1024 * 10 //10Mo is defined as max processing capability
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var contents []string
	for scanner.Scan() {
		line := Normalize(scanner.Text()) //Normalize line
		words := strings.Split(line, ";")
		for _, word := range words {
			contents = append(contents, strings.TrimSpace(word))
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	newFile.Content = contents
	return newFile
}
