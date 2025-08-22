package model

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*TreeElement
}
