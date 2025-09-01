package cmd

import "fmt"

func ShowHelp() {
	fmt.Println("Usage:")
	fmt.Println("  gosearch [options] <word>             Search for a word in the current directory")
	fmt.Println("  gosearch [options] -use <word> <name> Search for a word using a saved index")
	fmt.Println("  gosearch -save [path]                 Save an index of the current or specified path")
	fmt.Println("  gosearch -version                     Display version information")
	fmt.Println("  gosearch -help                        Display this help message")
	fmt.Println()

	fmt.Println("Options:")
	fmt.Println("  -v, --v, -version, --version          Show version information")
	fmt.Println("  -h, --h, -help, --help                Show usage help")
	fmt.Println("  -s, --s, -save, --save                Save index of current or specified path in Desktop/utils/index")
	fmt.Println("  -u, --u, -use, --use                  Use a specific saved index for search")
	fmt.Println()

	fmt.Println("Examples:")
	fmt.Println("  gosearch hello                        Search for 'hello' in current directory")
	fmt.Println("  gosearch -s                           Save index of current directory")
	fmt.Println("  gosearch -s /path/to/folder           Save index of specified folder")
	fmt.Println("  gosearch -u hello temp                Search for 'hello' in index named 'temp'")
	fmt.Println("  gosearch -version                     Show version info")
	fmt.Println("  gosearch -help                        Show this help message")
}
