package cmd

import "fmt"

func ShowVersion() {
	fmt.Printf("Version: %s\nDate: %s\nGo: %s\n", "1.2.0", "   16.09.2025", "     go 1.25.0 windows/amd64")
}
