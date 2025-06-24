package treebuilder

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type TreeElement struct {
	Name     string
	Path     string
	IsDir    bool
	Children []TreeElement
}

func GetFileTree(selected []string) []string {
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
			validFiles = append(validFiles, files...)
		} else if isValidFile(element) {
			validFiles = append(validFiles, element)
		}

		treeRoot.Children = append(treeRoot.Children, treeElement)
	}
	return validFiles
}

func isSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

func isExcludedPath(path string) bool {
	currentUser, err := user.Current()
	if err != nil {
		return false
	}

	excluded := []string{
		"Application Data",
		"$Recycle.Bin",
		"System Volume Information",
		"AppData",
		fmt.Sprintf("%s\\Cookies", currentUser.HomeDir),
		fmt.Sprintf("%s\\Documents\\Ma musique", currentUser.HomeDir),
		fmt.Sprintf("%s\\Documents\\Mes images", currentUser.HomeDir),
		fmt.Sprintf("%s\\Documents\\Mes vidéos", currentUser.HomeDir),
		fmt.Sprintf("%s\\Local Settings", currentUser.HomeDir),
		fmt.Sprintf("%s\\Menu Démarrer", currentUser.HomeDir),
		fmt.Sprintf("%s\\Mes documents", currentUser.HomeDir),
		fmt.Sprintf("%s\\Modèles", currentUser.HomeDir),
		fmt.Sprintf("%s\\NTUSER.DAT", currentUser.HomeDir),
		fmt.Sprintf("%s\\Recent", currentUser.HomeDir),
		fmt.Sprintf("%s\\SendTo", currentUser.HomeDir),
		fmt.Sprintf("%s\\Voisinage d'impression", currentUser.HomeDir),
		fmt.Sprintf("%s\\Voisinage réseau", currentUser.HomeDir),
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
			//log.Printf("%s est exclu ou symlink", childPath)
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
			validFiles = append(validFiles, subFiles...)
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
		e == ".png" || e == ".otb" || e == ".cff" || e == ".ttc" || e == ".base64" || e == ".syso")
}
