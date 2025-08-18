package cmd

import (
	"log"
	"os"
	"path/filepath"
	"search/utils"
)

func Indexate(path string) {
	//Step 1 - Get folder path
	var pathSelected string
	var err error
	var dir string

	if path == "-ui" { //if ui wanted
		pathSelected, err = utils.UISelectFolder()
	} else {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		pathSelected = dir + "\\" + path
	}

	//Step 2 - create file tree and filter it
	utils.CreateFileTree(pathSelected)
}
