package cmd

import "fmt"

func ShowHelp() {
	fmt.Printf("Usage:\n\tcmd <command> [<args>]\n\n")
	fmt.Printf("The commands are:\n")
	fmt.Printf("\tversion\t\tdisplay the version\n")
	fmt.Printf("\thelp\t\tdisplay this help\n")
	fmt.Printf("\t<word>\t\tsearch the word in the index generated from current position\n")
}
