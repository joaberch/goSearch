package utils

import (
	"github.com/ncruces/zenity"
	"log"
)

// UISelectFolder Select one or multiple folder
func UISelectFolder() string {
	file, err := zenity.SelectFile(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}
	return file
}
