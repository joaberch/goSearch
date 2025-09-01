package utils

import "github.com/joaberch/goSearch/internal/model"

// FlattenTree returns a flat list of all file elements from a TreeElement structure.
func FlattenTree(tree *model.TreeElement) []*model.TreeElement {
	var files []*model.TreeElement
	if !tree.IsDir {
		files = append(files, tree)
	} else {
		for _, child := range tree.Children {
			files = append(files, FlattenTree(child)...)
		}
	}
	return files
}
