package model

// FileData represents the path and content of a file.
type FileData struct { //FUTURE : GoDoc?
	Path    string
	Content map[int]string //Line x has .., .. and ..
}
