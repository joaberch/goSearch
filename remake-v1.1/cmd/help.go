package cmd

import "fmt"

func ShowHelp() {
	fmt.Printf("Usage:\n\tcmd <command> [<args>]\n\n")
	fmt.Printf("The commands are:\n")
	fmt.Printf("\tversion\tdisplay the version")
	fmt.Printf("\thelp\tdisplay this help")
	fmt.Printf("\tsearch [arg]\tsearch the argument in the index")
	fmt.Printf("\tindexate\tcreate an index from the path given")
	fmt.Printf("\tdisplay\tdisplay the structure of the file indexed")
}
