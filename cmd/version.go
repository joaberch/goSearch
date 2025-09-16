package cmd

import "fmt"

// ShowVersion prints the program version, build date, and Go runtime/target on standard output.
// It emits three lines prefixed with "Version:", "Date:", and "Go:" respectively.
func ShowVersion() {
	fmt.Printf("Version: %s\nDate: %s\nGo: %s\n", "1.2.0", "   16.09.2025", "     go 1.25.0 windows/amd64")
}
