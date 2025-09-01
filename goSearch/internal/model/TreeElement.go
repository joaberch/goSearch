package model

// TreeElement represents a node in a directory tree.
type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*TreeElement
}
