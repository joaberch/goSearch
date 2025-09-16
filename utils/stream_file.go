package utils

import (
	"bufio"
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
	"strings"
)

// StreamFile reads the content of the file specified by the given TreeElement
// and returns it as a FileData object. Each line is processed by Normalize()
// before being stored in the Content map, where keys represent line numbers
// and values are the normalized line strings.
//
// Behavior:
//   - If the file cannot be opened, an empty FileData is returned (with only Path set).
//   - The scanner is configured with a maximum buffer capacity of 10 MB to handle long lines.
//   - Any errors during scanning or closing the file cause the program to terminate via log.Fatal.
//
// Parameters:
//
//	element *model.TreeElement: A struct containing the path to the target file.
//
// Returns:
//
//	model.FileData: A struct containing the file path and a map of line numbers to normalized lines.
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
