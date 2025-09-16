package cmd

import "fmt"

// ShowHelp prints the command-line usage, commands, options, examples, and notes for the `gosearch` tool to standard output.
// The message is static and formatted; the function performs no parsing and does not return an error.
func ShowHelp() {
	fmt.Println("Usage:")
	fmt.Println("  gosearch [options] <search_term>")
	fmt.Println()

	fmt.Println("Commands:")
	fmt.Println("  -h, --help                			Show this help message")
	fmt.Println("  -v, --version          				Show version information")
	fmt.Println("  -s, --save                			Save an index of the given folder (defaults to current directory if not specified)")
	fmt.Println("  -u, --use                  			Use a specific saved index for search")
	fmt.Println()

	fmt.Println("Options:")
	fmt.Println("  -m, --match <mode>					Set match mode: 'contains' (default), 'exact' or 'regex'")

	fmt.Println("Examples:")
	fmt.Println("  gosearch hello                       Search for 'hello' in current directory")
	fmt.Println("  gosearch -s                          Save index of current directory")
	fmt.Println("  gosearch -s /path/to/folder          Save index of specified folder")
	fmt.Println("  gosearch -u temp hello              	Search for 'hello' in an index named 'temp', loose verification")
	fmt.Println("  gosearch hello -m exact -u temp		Search for 'hello' in an index named 'temp', strict verification")
	fmt.Println()

	fmt.Println("Notes:")
	fmt.Println("  - The search is case-insensitive")
	fmt.Println("  - The options and commands can be placed before or after the search term")
}
