package treebuilder

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []TreeElement
}

type TreeResult struct {
	Tree       TreeElement
	ValidFiles []string
}

func GetFileTree(selected []string) TreeResult {
	var validFiles []string

	treeRoot := TreeElement{
		Name:     "root",
		Path:     "",
		IsDir:    true,
		Children: []TreeElement{},
	}

	for _, element := range selected {
		info, err := os.Stat(element)
		if err != nil {
			log.Println(err)
			continue
		}

		treeElement := TreeElement{
			Name:  info.Name(),
			Path:  element,
			IsDir: info.IsDir(),
		}
		if info.IsDir() {
			children, files := recursiveTree(element)
			treeElement.Children = children
			validFiles = append(validFiles, files...) //TODO - understand
		} else if isValidFile(element) {
			validFiles = append(validFiles, element)
		}

		treeRoot.Children = append(treeRoot.Children, treeElement)
	}
	return TreeResult{Tree: treeRoot, ValidFiles: validFiles}
}

func isSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

func isExcludedPath(path string) bool {
	excluded := []string{
		"Application Data",
		"$Recycle.Bin",
		"System Volume Information",
		"AppData",
		"px59nyu\\Cookies",
		"Documents\\Ma musique",
		"Documents\\Mes images",
		"Documents\\Mes vidéos",
		"px59nyu\\Local Settings",
		"px59nyu\\Menu Démarrer",
		"px59nyu\\Mes documents",
		"px59nyu\\Modèles",
		"px59nyu\\NTUSER.DAT",
		"OneDrive - Education Vaud",
		"px59nyu\\Recent",
		"px59nyu\\SendTo",
		"px59nyu\\Voisinage d'impression",
		"px59nyu\\Voisinage réseau",
	}
	path = strings.ToLower(path)
	for _, ex := range excluded {
		if strings.Contains(path, strings.ToLower(ex)) {
			return true
		}
	}
	return false
}

func recursiveTree(path string) ([]TreeElement, []string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("Impossible d'accéder à %s : %v", path, err) //If it requires admin right we just stop
		return []TreeElement{}, []string{}
	}

	var children []TreeElement
	var validFiles []string

	for _, entry := range entries {
		childPath := path + "\\" + entry.Name()

		if isExcludedPath(childPath) || isSymlink(childPath) {
			log.Printf("%s est exclu ou symlink", childPath)
			continue
		}

		child := TreeElement{
			Name:  entry.Name(),
			Path:  childPath,
			IsDir: entry.IsDir(),
		}

		if entry.IsDir() {
			subChildren, subFiles := recursiveTree(childPath)
			child.Children = subChildren
			validFiles = append(validFiles, subFiles...) //TODO - understand
		} else if isValidFile(childPath) {
			validFiles = append(validFiles, childPath)
		}
		children = append(children, child)
	}
	return children, validFiles
}

func isValidFile(path string) bool {
	e := strings.ToLower(filepath.Ext(path)) //e = extension
	return !(e == ".bin" || e == ".exe" || e == ".dll" || e == ".iso" || e == ".lnk" || e == ".mp4" || e == ".zip" || e == ".ttf" || e == ".otf" ||
		e == ".png" || e == ".otb" || e == ".cff" || e == ".ttc" || e == ".base64" || e == ".syso") //TODO - add other and optimize
}
