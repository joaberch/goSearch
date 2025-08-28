package utils

import (
	"goSearch/internal/model"
)

// CreateIndex creates an inverted index
func CreateIndex(files []model.FileData) model.InvertedIndex {
	index := make(model.InvertedIndex) //Instantiate my map
	for _, file := range files {       //Process each file one by one - TODO: Should I create packages of 5 by 5 for goroutine?
		seen := make(map[string]bool)        //"myword": true||false //if has already been processed
		for _, token := range file.Content { //Process each word - TODO: Should I create packages of x by x for goroutine?
			if !seen[token] { //If hasn't already been processed - TODO: Should I remove it and add a weight system?
				index[token] = append(index[token], file.Path) //Add file path
				seen[token] = true                             //weight system?
			}
		}
	}
	return index
}
