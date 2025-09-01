package utils

import (
	"github.com/joaberch/goSearch/internal/model"
)

// Indexate builds an inverted index from all valid files found in the given directory path.
func Indexate(path string) model.InvertedIndex {
	tree := CreateFileTree(path)

	var contents []model.FileData
	for _, file := range FlattenTree(&tree) {
		streamRes := StreamFile(file)
		contents = append(contents, streamRes)
	}

	return CreateIndex(contents)
}
