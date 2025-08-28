package utils

import "goSearch/internal/model"

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
