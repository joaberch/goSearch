package model

import "encoding/xml"

type IndexDocument struct { //The whole XML file
	XMLName xml.Name     `xml:"index"`
	Entries []IndexEntry `xml:"entry"`
}
