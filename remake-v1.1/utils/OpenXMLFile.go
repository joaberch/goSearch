package utils

import (
	"encoding/xml"
	"os"
)

func OpenXMLFile(path string) (*xml.Decoder, *os.File, error) {
	source, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return xml.NewDecoder(source), source, err
}
