package utils

import (
	"search/internal/model"
)

func Indexate(path string) model.InvertedIndex {
	//Step 1 - create file tree and filter it (filter is inside CreateFileTree)
	tree := CreateFileTree(path)

	//Step 2 - Stream files (and normalize it in StreamFile)
	var contents []model.FileData
	for _, file := range FlattenTree(&tree) {
		streamRes := StreamFile(file)
		contents = append(contents, streamRes)
	}

	//Step 3 - Save result in inverted index var
	return CreateIndex(contents)
}
