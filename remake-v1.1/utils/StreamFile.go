package utils

import (
	"bufio"
	"os"
	"search/internal/model"
	"strings"
)

func StreamFile(element *model.TreeElement) model.FileData {
	newFile := model.FileData{
		Path: element.Path,
	}

	file, err := os.Open(newFile.Path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

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
	newFile.Content = contents
	return newFile
}
