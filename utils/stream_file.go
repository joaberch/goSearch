package utils

import (
	"bufio"
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
	"strings"
)

// StreamFile reads a file line by line, normalizes its content, and returns a FileData structure containing the extracted words.
func StreamFile(element *model.TreeElement) model.FileData {
	newFile := model.FileData{
		Path:    element.Path,
		Content: make(map[int]string),
	}

	file, err := os.Open(newFile.Path)
	if err != nil {
		log.Printf("Error opening file %s : %v", newFile.Path, err)
		return newFile
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)    // Error if the line length is longer than 65'536 character so we define a max capacity
	const maxCapacity = 1024 * 1024 * 10 //10Mo is defined as max processing capability
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	lineNumber := 1
	for scanner.Scan() {
		normalizedWords := Normalize(scanner.Text()) //Normalize line
		newFile.Content[lineNumber] = strings.Join(normalizedWords, " ")
		lineNumber++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return newFile
}
