package model

import "encoding/xml"

// IndexDocument represents the root of an XML index file.
type IndexDocument struct { //The whole XML file
	XMLName xml.Name     `xml:"index"`
	Entries []IndexEntry `xml:"entry"`
}
