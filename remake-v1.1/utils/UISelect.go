package utils

import "github.com/ncruces/zenity"

// UISelectFolder opens a directory picker and returns the selected path.
// It returns a non-nil error if the user cancels the dialog or if an error occurs.
func UISelectFolder() (string, error) {
	file, err := zenity.SelectFile(zenity.Directory())
	if err != nil {
		return "", err
	}
	return file, nil
}
