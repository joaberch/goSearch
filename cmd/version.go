package cmd

import "fmt"

func ShowVersion() {
	fmt.Printf("Version: %s\nDate: %s\nGo: %s\n", "1.1.0", "   05.09.2025", "     go 1.25.0 windows/amd64")
}
